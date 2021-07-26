package middlewares

import (
	"OrderService/internal/database"
	"context"
	"time"

	//import the service package
	"OrderService/internal/service"
	"github.com/go-kit/kit/log"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   service.Service
}

func (l *LoggingMiddleware) GetByID(ctx context.Context, id uint) (order database.OrderOut,err error) {
	defer func(begin time.Time) {
		err := l.Logger.Log("method", "get by id", "id",id ,"name", order.FirstName+" "+order.LastName,"err", err, "took", time.Since(begin))
		if err != nil {
			return
		}
	}(time.Now())
	order,err=l.Next.GetByID(ctx,id)
	return
}

func (l *LoggingMiddleware) Search(ctx context.Context, search string, startDate time.Time, endDate time.Time) (orders []database.OrderOut,err error) {
	defer func(begin time.Time) {
		err := l.Logger.Log("method", "search", "orders", len(orders),"err", err, "took", time.Since(begin))
		if err != nil {
			return
		}
	}(time.Now())
	orders,err=l.Next.Search(ctx,search,startDate,endDate)
	return
}

func (l *LoggingMiddleware) Create(ctx context.Context, data database.OrderIn) (msg string,err error) {
	defer func(begin time.Time) {
		err := l.Logger.Log("method", "create", "message", msg,"err", err, "took", time.Since(begin))
		if err != nil {
			return
		}
	}(time.Now())
	msg,err=l.Next.Create(ctx,data)
	return
}

func (l *LoggingMiddleware) Delete(ctx context.Context, id uint) (msg string,err error) {
	defer func(begin time.Time) {
		err := l.Logger.Log("method", "delete", "id",id ,"message", msg,"err", err, "took", time.Since(begin))
		if err != nil {
			return
		}
	}(time.Now())
	msg,err=l.Next.Delete(ctx,id)
	return
}

func (l *LoggingMiddleware) Total(ctx context.Context) (total float32,err error) {
	defer func(begin time.Time) {
		err := l.Logger.Log("method", "total", "err", err, "took", time.Since(begin))
		if err != nil {
			return
		}
	}(time.Now())
	total,err=l.Next.Total(ctx)
	return
}

func (l *LoggingMiddleware) Top(ctx context.Context, count uint) (products []database.ProductOut,err error) {
	defer func(begin time.Time) {
		err := l.Logger.Log("method", "top", "products", len(products),"err", err, "took", time.Since(begin))
		if err != nil {
			return
		}
	}(time.Now())
	products,err=l.Next.Top(ctx,count)
	return
}


