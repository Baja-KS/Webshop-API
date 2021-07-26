package middlewares

import (
	"OrderService/internal/database"
	"OrderService/internal/service"
	"context"
	"github.com/go-kit/kit/metrics"
	"time"
)

type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	Next           service.Service
}

//TODO metrics

func (i *InstrumentingMiddleware) GetByID(ctx context.Context, id uint) (database.OrderOut, error) {
	panic("implement me")
}

func (i *InstrumentingMiddleware) Search(ctx context.Context, search string, startDate time.Time, endDate time.Time) ([]database.OrderOut, error) {
	panic("implement me")
}

func (i *InstrumentingMiddleware) Create(ctx context.Context, data database.OrderIn) (string, error) {
	panic("implement me")
}

func (i *InstrumentingMiddleware) Delete(ctx context.Context, id uint) (string, error) {
	panic("implement me")
}

func (i *InstrumentingMiddleware) Total(ctx context.Context) (float32, error) {
	panic("implement me")
}

func (i *InstrumentingMiddleware) Top(ctx context.Context, count uint) ([]database.ProductOut, error) {
	panic("implement me")
}
