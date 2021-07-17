package middlewares

import (
	"CategoryService/internal/database"
	"context"
	"time"

	//import the service package
	"CategoryService/internal/service"
	"github.com/go-kit/kit/log"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   service.Service
}

func (l *LoggingMiddleware) GetAll(ctx context.Context) (categories []database.CategoryOut,err error) {
	defer func(begin time.Time) {
		err := l.Logger.Log("method", "getall", "categories", len(categories),"err", err, "took", time.Since(begin))
		if err != nil {
			return
		}
	}(time.Now())
	categories,err=l.Next.GetAll(ctx)
	return
}

func (l *LoggingMiddleware) Create(ctx context.Context, data database.CategoryIn) (msg string,err error) {
	defer func(begin time.Time) {
		err := l.Logger.Log("method", "create", "message", msg,"err", err, "took", time.Since(begin))
		if err != nil {
			return
		}
	}(time.Now())
	msg,err=l.Next.Create(ctx,data)
	return
}

func (l *LoggingMiddleware) Update(ctx context.Context, id uint, data database.CategoryIn) (msg string,err error) {
	defer func(begin time.Time) {
		err := l.Logger.Log("method", "update", "id",id ,"message", msg,"err", err, "took", time.Since(begin))
		if err != nil {
			return
		}
	}(time.Now())
	msg,err=l.Next.Update(ctx,id,data)
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

func (l *LoggingMiddleware) Products(ctx context.Context, id uint) (products []database.ProductOut,err error) {
	defer func(begin time.Time) {
		err := l.Logger.Log("method", "products", "products", len(products),"err", err, "took", time.Since(begin))
		if err != nil {
			return
		}
	}(time.Now())
	products,err=l.Next.Products(ctx,id)
	return
}

func (l *LoggingMiddleware) GetByID(ctx context.Context, id uint) (category database.CategoryOut,err error) {
	defer func(begin time.Time) {
		err := l.Logger.Log("method", "get by id", "id",id ,"name", category.Name,"err", err, "took", time.Since(begin))
		if err != nil {
			return
		}
	}(time.Now())
	category,err=l.Next.GetByID(ctx,id)
	return
}

func (l *LoggingMiddleware) GetByGroupID(ctx context.Context, id uint) (categories []database.CategoryOut,err error) {
	defer func(begin time.Time) {
		err := l.Logger.Log("method", "get by group", "categories", len(categories),"err", err, "took", time.Since(begin))
		if err != nil {
			return
		}
	}(time.Now())
	categories,err=l.Next.GetByGroupID(ctx,id)
	return
}


