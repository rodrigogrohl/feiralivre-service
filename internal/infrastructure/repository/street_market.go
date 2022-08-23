package repository

import (
	"context"
	"strings"
	"sync"

	"github.com/rodrigogrohl/feiralivre-service/internal/infrastructure/repository/database"
	"github.com/rodrigogrohl/feiralivre-service/pkg/canonical"
	"github.com/sirupsen/logrus"
)

var (
	_streetMarketRepository StreetMarketRepository

	// Assure Singleton into repos initialize
	_onceStreetMarketRepository sync.Once
)

type StreetMarketRepository interface {
	Get(ctx context.Context, id int64) (*canonical.StreetMarket, error)
	Insert(ctx context.Context, sm *canonical.StreetMarket) (int64, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, sm *canonical.StreetMarket) error
	QueryBy(ctx context.Context, filter *canonical.StreetMarketFilter) ([]*canonical.StreetMarket, error)
}

type streetMarketRepositoryImpl struct {
	db *database.DBStruct
}

func InitStreetMarketRepository() StreetMarketRepository {
	_onceStreetMarketRepository.Do(func() {
		_streetMarketRepository = &streetMarketRepositoryImpl{
			db: database.DatabaseConnect(),
		}
	})
	return _streetMarketRepository
}

func (r *streetMarketRepositoryImpl) Get(ctx context.Context, id int64) (*canonical.StreetMarket, error) {
	logrus.WithField("id", id).Info("get StreetMarket")
	row, err := r.db.Conn.Query("SELECT * FROM street_market WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	sm := &canonical.StreetMarket{}
	if row.Next() {
		err = row.Scan(&sm.Id, &sm.Longitude, &sm.Latitude, &sm.SectorCense, &sm.AreaPonderate, &sm.DistrictCode, &sm.District, &sm.SubTownCode, &sm.SubTown, &sm.Region5, &sm.Region8, &sm.Name, &sm.Registry, &sm.Address, &sm.Number, &sm.Neighborhood, &sm.Reference)
		if err != nil {
			return nil, err
		}
		return sm, err
	}
	return nil, canonical.DBEntityNotFound
}

func (r *streetMarketRepositoryImpl) Insert(ctx context.Context, sm *canonical.StreetMarket) (int64, error) {
	logrus.WithField("street_market", sm).Info("inserting")
	stmt, err := r.db.Conn.Prepare("INSERT INTO street_market (id,longitude,latitude,sector_cense,area_ponderate,district_code,district,sub_town_code,sub_town,region5,region8,name_alias,registry,addr,addr_number,neighborhood,reference) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var refId *int64
	if sm.Id > 0 {
		refId = &sm.Id
	}

	res, err := stmt.Exec(refId, sm.Longitude, sm.Latitude, sm.SectorCense, sm.AreaPonderate, sm.DistrictCode, sm.District, sm.SubTownCode, sm.SubTown, sm.Region5, sm.Region8, sm.Name, sm.Registry, sm.Address, sm.Number, sm.Neighborhood, sm.Reference)
	defer stmt.Close()

	if err != nil {
		return 0, err
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *streetMarketRepositoryImpl) Delete(ctx context.Context, id int64) error {
	logrus.WithField("id", id).Info("delete StreetMarket")
	_, err := r.db.Conn.Exec("DELETE FROM street_market WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *streetMarketRepositoryImpl) Update(ctx context.Context, sm *canonical.StreetMarket) error {
	logrus.WithField("street_market", sm).Info("updating")
	stmt, err := r.db.Conn.Prepare("UPDATE street_market SET longitude = ?,latitude = ?, sector_cense = ?, area_ponderate = ?, district_code = ?, district = ?, sub_town_code = ?, sub_town = ?, region5 = ?, region8 = ?, name_alias = ?, registry = ?, addr = ?, addr_number = ?, neighborhood = ?, reference = ? WHERE id = ?;")
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(sm.Longitude, sm.Latitude, sm.SectorCense, sm.AreaPonderate, sm.DistrictCode, sm.District, sm.SubTownCode, sm.SubTown, sm.Region5, sm.Region8, sm.Name, sm.Registry, sm.Address, sm.Number, sm.Neighborhood, sm.Reference, sm.Id)

	if err != nil {
		return err
	}

	return nil
}

func (r *streetMarketRepositoryImpl) QueryBy(ctx context.Context, filter *canonical.StreetMarketFilter) ([]*canonical.StreetMarket, error) {
	whereCol, whereVal := filter.Extract()
	if len(whereCol) < 1 {
		return nil, canonical.MissingFieldsError
	}

	queryStr := "SELECT * FROM street_market WHERE " + strings.Join(whereCol, " AND ")
	rows, err := r.db.Conn.Query(queryStr, whereVal...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := []*canonical.StreetMarket{}
	for rows.Next() {
		sm := canonical.StreetMarket{}
		err = rows.Scan(&sm.Id, &sm.Longitude, &sm.Latitude, &sm.SectorCense, &sm.AreaPonderate, &sm.DistrictCode, &sm.District, &sm.SubTownCode, &sm.SubTown, &sm.Region5, &sm.Region8, &sm.Name, &sm.Registry, &sm.Address, &sm.Number, &sm.Neighborhood, &sm.Reference)
		if err != nil {
			return nil, err
		}
		data = append(data, &sm)
	}
	return data, nil
}
