package application

import (
	"context"
	"sync"

	"github.com/rodrigogrohl/feiralivre-service/internal/infrastructure/repository"
	"github.com/rodrigogrohl/feiralivre-service/internal/presentation/metrics"
	"github.com/rodrigogrohl/feiralivre-service/pkg/canonical"
)

var (
	_streetMarketService     StreetMarketService
	_onceStreetMarketService sync.Once
)

type StreetMarketService interface {
	Get(ctx context.Context, id int64) (*canonical.StreetMarket, error)
	Create(ctx context.Context, sm *canonical.StreetMarket) error
	Remove(ctx context.Context, id int64) error
	Update(ctx context.Context, sm *canonical.StreetMarket) error
	QueryBy(ctx context.Context, filter *canonical.StreetMarketFilter) ([]*canonical.StreetMarket, error)
}

type streetMarketServiceImpl struct {
	repos repository.StreetMarketRepository
}

func InitStreetMarketService() StreetMarketService {
	_onceStreetMarketService.Do(func() {
		_streetMarketService = &streetMarketServiceImpl{
			repos: repository.InitStreetMarketRepository(),
		}
	})
	return _streetMarketService
}

func (s *streetMarketServiceImpl) Get(ctx context.Context, id int64) (*canonical.StreetMarket, error) {
	sm, err := s.repos.Get(ctx, id)
	defer metrics.AddCounterErr("street_market", "Get", err)
	return sm, err
}

func (s *streetMarketServiceImpl) Create(ctx context.Context, sm *canonical.StreetMarket) error {
	id, err := s.repos.Insert(ctx, sm)
	defer metrics.AddCounterErr("street_market", "Create", err)
	if err != nil {
		return err
	}
	sm.Id = id
	return nil
}

func (s *streetMarketServiceImpl) Remove(ctx context.Context, id int64) error {
	err := s.repos.Delete(ctx, id)
	defer metrics.AddCounterErr("street_market", "Remove", err)
	return err
}

func (s *streetMarketServiceImpl) Update(ctx context.Context, sm *canonical.StreetMarket) error {
	err := s.repos.Update(ctx, sm)
	defer metrics.AddCounterErr("street_market", "Update", err)
	return err
}

func (s *streetMarketServiceImpl) QueryBy(ctx context.Context, filter *canonical.StreetMarketFilter) ([]*canonical.StreetMarket, error) {
	resultList, err := s.repos.QueryBy(ctx, filter)
	defer metrics.AddCounterErr("street_market", "QueryBy", err)
	return resultList, err
}
