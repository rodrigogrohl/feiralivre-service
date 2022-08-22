package application

import (
	"context"

	"github.com/rodrigogrohl/feiralivre-service/internal/application/canonical"
)

type StreetMarketService interface {
	Create(ctx context.Context, sm canonical.StreetMarket) (int64, error)
	Remove(ctx context.Context, id int64) error
	Update(ctx context.Context, sm canonical.StreetMarket) error
	QueryBy(ctx context.Context, filter canonical.StreetMarketFilter) ([]*canonical.StreetMarket, error)
}
