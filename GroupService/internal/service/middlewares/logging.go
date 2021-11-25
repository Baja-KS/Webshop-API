package middlewares

import (
	"GroupService/internal/database"
	"context"
	"time"

	//import the service package
	"GroupService/internal/service"
	"github.com/go-kit/kit/log"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   service.Service
}

func (l *LoggingMiddleware) GetAll(ctx context.Context) (groups []database.GroupOutWithCategories,err error) {
	defer func(begin time.Time) {
		err := l.Logger.Log("method", "getall", "groups", len(groups),"err", err, "took", time.Since(begin))
		if err != nil {
			return
		}
	}(time.Now())
	groups,err=l.Next.GetAll(ctx)
	return
}

func (l *LoggingMiddleware) Create(ctx context.Context, data database.GroupIn) (msg string,err error) {
	defer func(begin time.Time) {
		err := l.Logger.Log("method", "create", "message", msg,"err", err, "took", time.Since(begin))
		if err != nil {
			return
		}
	}(time.Now())
	msg,err=l.Next.Create(ctx,data)
	return
}

func (l *LoggingMiddleware) Update(ctx context.Context, id uint, data database.GroupIn) (msg string,err error) {
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

func (l *LoggingMiddleware) Categories(ctx context.Context, id uint) (categories []database.CategoryOut,err error) {
	defer func(begin time.Time) {
		err := l.Logger.Log("method", "categories", "group id",id,"categories", len(categories),"err", err, "took", time.Since(begin))
		if err != nil {
			return
		}
	}(time.Now())
	categories,err=l.Next.Categories(ctx,id)
	return
}

func (l *LoggingMiddleware) GetByID(ctx context.Context, id uint) (group database.GroupOut,err error) {
	defer func(begin time.Time) {
		err := l.Logger.Log("method", "get by id", "id",id ,"name", group.Name,"err", err, "took", time.Since(begin))
		if err != nil {
			return
		}
	}(time.Now())
	group,err=l.Next.GetByID(ctx,id)
	return
}
