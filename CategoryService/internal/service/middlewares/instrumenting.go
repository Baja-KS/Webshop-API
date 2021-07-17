package middlewares

import (
	"CategoryService/internal/database"
	"CategoryService/internal/service"
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

func (i *InstrumentingMiddleware) GetAll(ctx context.Context) (categories []database.CategoryOut,err error) {
	defer func(begin time.Time) {
		lvs:=[]string{"method","GetAll","category_id", "none","error",fmt.Sprint(err!=nil)}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	categories,err=i.Next.GetAll(ctx)
	return
}

func (i *InstrumentingMiddleware) Create(ctx context.Context, data database.CategoryIn) (msg string,err error) {
	defer func(begin time.Time) {
		lvs:=[]string{"method","Create","category_id", "none","error",fmt.Sprint(err!=nil)}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	msg,err=i.Next.Create(ctx,data)
	return
}

func (i *InstrumentingMiddleware) Update(ctx context.Context, id uint, data database.CategoryIn) (msg string,err error) {
	defer func(begin time.Time) {
		lvs:=[]string{"method","Update","category_id", strconv.Itoa(int(id)),"error",fmt.Sprint(err!=nil)}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	msg,err=i.Next.Update(ctx,id,data)
	return
}

func (i *InstrumentingMiddleware) Delete(ctx context.Context, id uint) (msg string,err error) {
	defer func(begin time.Time) {
		lvs:=[]string{"method","Delete","category_id", strconv.Itoa(int(id)),"error",fmt.Sprint(err!=nil)}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	msg,err=i.Next.Delete(ctx,id)
	return
}

func (i *InstrumentingMiddleware) Products(ctx context.Context, id uint) (products []database.ProductOut,err error) {
	defer func(begin time.Time) {
		lvs:=[]string{"method","Products","category_id", strconv.Itoa(int(id)),"error",fmt.Sprint(err!=nil)}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	products,err=i.Next.Products(ctx,id)
	return
}

func (i *InstrumentingMiddleware) GetByID(ctx context.Context, id uint) (category database.CategoryOut,err error) {
	defer func(begin time.Time) {
		lvs:=[]string{"method","GetByID","category_id", strconv.Itoa(int(id)),"error",fmt.Sprint(err!=nil)}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	category,err=i.Next.GetByID(ctx,id)
	return
}

func (i *InstrumentingMiddleware) GetByGroupID(ctx context.Context, id uint) (categories []database.CategoryOut,err error) {
	defer func(begin time.Time) {
		lvs:=[]string{"method","GetByGroupID","category_id", "none","error",fmt.Sprint(err!=nil)}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	categories,err=i.Next.GetByGroupID(ctx,id)
	return
}
