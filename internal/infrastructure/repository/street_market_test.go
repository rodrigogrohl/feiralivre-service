package repository

import (
	"context"
	"testing"

	"github.com/rodrigogrohl/feiralivre-service/internal/infrastructure/config"
	"github.com/rodrigogrohl/feiralivre-service/pkg/canonical"
	"github.com/stretchr/testify/assert"
)

func TestInsertStreetMarket_OK(t *testing.T) {
	t.Skip()
	config.InitializeTest()
	repos := InitStreetMarketRepository()
	id, err := repos.Insert(context.Background(), &canonical.StreetMarket{
		Longitude:     99999,
		Latitude:      -9999999,
		SectorCense:   3550300091,
		AreaPonderate: 35503080,
		DistrictCode:  10,
		District:      "MOEMA",
		SubTownCode:   99,
		SubTown:       "VILA MARIANA",
		Region5:       "Sul",
		Region8:       "Sul 1",
		Name:          "IBIRAPUERA",
		Registry:      "3009-0",
		Address:       "RUA JOINVILE",
		Number:        "323",
		Neighborhood:  "IBIRAPUERA",
		Reference:     "S/INFOR",
	})
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, id, int64(0))
}

func TestGetStreetMarket_OK(t *testing.T) {
	config.InitializeTest()
	repos := InitStreetMarketRepository()
	sm, err := repos.Get(context.Background(), 1)
	assert.Nil(t, err)
	assert.NotNil(t, sm)
}

func TestGetStreetMarket_NOK(t *testing.T) {
	config.InitializeTest()
	repos := InitStreetMarketRepository()
	sm, err := repos.Get(context.Background(), 999999999999)
	assert.NotNil(t, err)
	assert.Nil(t, sm)
}

func TestQueryByName_OK(t *testing.T) {
	config.InitializeTest()
	repos := InitStreetMarketRepository()
	list, err := repos.QueryBy(context.Background(), &canonical.StreetMarketFilter{
		Name: "VILA",
	})
	assert.Nil(t, err)
	assert.NotNil(t, list)
	assert.Greater(t, len(list), 0)
}

func TestQueryByRegion5_OK(t *testing.T) {
	config.InitializeTest()
	repos := InitStreetMarketRepository()
	list, err := repos.QueryBy(context.Background(), &canonical.StreetMarketFilter{
		Region5: "Sul",
	})
	assert.Nil(t, err)
	assert.NotNil(t, list)
	assert.Greater(t, len(list), 0)
}

func TestQueryByNeighborhood_OK(t *testing.T) {
	config.InitializeTest()
	repos := InitStreetMarketRepository()
	list, err := repos.QueryBy(context.Background(), &canonical.StreetMarketFilter{
		Neighborhood: "IBIRA",
	})
	assert.Nil(t, err)
	assert.NotNil(t, list)
	assert.Greater(t, len(list), 0)
}

func TestQueryByDistrict_OK(t *testing.T) {
	config.InitializeTest()
	repos := InitStreetMarketRepository()
	list, err := repos.QueryBy(context.Background(), &canonical.StreetMarketFilter{
		District: "MO",
	})
	assert.Nil(t, err)
	assert.NotNil(t, list)
	assert.Greater(t, len(list), 0)
}
func TestQueryBy_NOK(t *testing.T) {
	config.InitializeTest()
	repos := InitStreetMarketRepository()
	list, err := repos.QueryBy(context.Background(), &canonical.StreetMarketFilter{})
	assert.NotNil(t, err)
	assert.Nil(t, list)
}

func TestDelete_OK(t *testing.T) {
	config.InitializeTest()
	repos := InitStreetMarketRepository()
	err := repos.Delete(context.Background(), 1)
	assert.Nil(t, err)
}

func TestUpdateStreetMarket_OK(t *testing.T) {
	config.InitializeTest()
	repos := InitStreetMarketRepository()
	sm, _ := repos.Get(context.Background(), 5)
	sm.Name = sm.Name + "-up"
	sm.District = sm.District + "-up"
	err := repos.Update(context.Background(), sm)
	assert.Nil(t, err)
	assert.NotNil(t, sm)
}
