package middlewares

import (
	"GroupService/internal/database"
	"GroupService/internal/service"
	"context"
	"fmt"
	"github.com/go-kit/kit/metrics"
	"strconv"
	"time"
)

type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	Next           service.Service
}

func (i *InstrumentingMiddleware) GetAll(ctx context.Context) (groups []database.GroupOutWithCategories,err error) {
	defer func(begin time.Time) {
		lvs:=[]string{"method","GetAll","group_id", "none","error",fmt.Sprint(err!=nil)}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	groups,err=i.Next.GetAll(ctx)
	return
}

func (i *InstrumentingMiddleware) Create(ctx context.Context, data database.GroupIn) (msg string,err error) {
	defer func(begin time.Time) {
		lvs:=[]string{"method","Create","group_id", "none","error",fmt.Sprint(err!=nil)}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	msg,err=i.Next.Create(ctx,data)
	return
}

func (i *InstrumentingMiddleware) Update(ctx context.Context, id uint, data database.GroupIn) (msg string,err error) {
	defer func(begin time.Time) {
		lvs:=[]string{"method","Update","group_id", strconv.Itoa(int(id)),"error",fmt.Sprint(err!=nil)}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	msg,err=i.Next.Update(ctx,id,data)
	return
}

func (i *InstrumentingMiddleware) Delete(ctx context.Context, id uint) (msg string,err error) {
	defer func(begin time.Time) {
		lvs:=[]string{"method","Delete","group_id", strconv.Itoa(int(id)),"error",fmt.Sprint(err!=nil)}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	msg,err=i.Next.Delete(ctx,id)
	return
}

func (i *InstrumentingMiddleware) Categories(ctx context.Context, id uint) (categories []database.CategoryOut,err error) {
	defer func(begin time.Time) {
		lvs:=[]string{"method","Categories","group_id", strconv.Itoa(int(id)),"error",fmt.Sprint(err!=nil)}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	categories,err=i.Next.Categories(ctx,id)
	return
}

func (i *InstrumentingMiddleware) GetByID(ctx context.Context, id uint) (group database.GroupOut,err error) {
	defer func(begin time.Time) {
		lvs:=[]string{"method","GetByID","group_id", strconv.Itoa(int(id)),"error",fmt.Sprint(err!=nil)}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	group,err=i.Next.GetByID(ctx,id)
	return
}
