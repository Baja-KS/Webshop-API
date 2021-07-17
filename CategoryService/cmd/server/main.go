package main

import (
	"CategoryService/internal/database"
	"CategoryService/internal/service"
	"CategoryService/internal/service/endpoints"
	"CategoryService/internal/service/middlewares"
	"CategoryService/internal/service/transport"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"net/http"
	"os"
)


func main() {
	logger:=log.NewLogfmtLogger(os.Stderr)
	var err error
	db,err:=database.NewDatabase()
	if err != nil {
		log.With(logger,"err",err)
	}
	err=database.Migrate(db)
	if err != nil {
		log.With(logger,"err",err)
	}

	requestCount:=kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace:   "bajaks_webshop_api",
		Subsystem:   "category_service",
		Name:        "request_count",
		Help:        "Number of request received",
	},[]string{"method","category_id","error"})
	requestLatency:=kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace:   "bajaks_webshop_api",
		Subsystem:   "category_service",
		Name:        "request_latency",
		Help:        "Total duration of requests in microseconds",
	},[]string{"method","category_id","error"})

	var svc service.Service
	svc=&service.CategoryService{DB: db}
	svc=&middlewares.AuthenticationMiddleware{Next: svc}
	svc=&middlewares.LoggingMiddleware{Logger: logger,Next: svc}
	svc=&middlewares.InstrumentingMiddleware{
		RequestCount:   requestCount,
		RequestLatency: requestLatency,
		Next:           svc,
	}

	ep:=endpoints.NewEndpointSet(svc)
	err = http.ListenAndServe(":8080", transport.NewHTTPHandler(ep))
	if err != nil {
		log.With(logger,"err",err)
	}
	log.With(logger,"msg","Listening to port")
}
