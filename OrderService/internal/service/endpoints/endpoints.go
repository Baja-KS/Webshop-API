package endpoints

import (
	"OrderService/internal/service"
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func ParseIDFromURL(r *http.Request) (uint, error) {
	params:=mux.Vars(r)
	idStr:=params["id"]
	id,err:=strconv.ParseUint(idStr,10,32)
	if err != nil {
		return 0,err
	}
	return uint(id),nil
}

type EndpointSet struct {
	GetByIDEndpoint endpoint.Endpoint
	SearchEndpoint endpoint.Endpoint
	CreateEndpoint endpoint.Endpoint
	DeleteEndpoint endpoint.Endpoint
	TotalEndpoint endpoint.Endpoint
	TopEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc service.Service) EndpointSet {
	return EndpointSet{
		GetByIDEndpoint:    MakeGetByIDEndpoint(svc),
		SearchEndpoint:    MakeSearchEndpoint(svc),
		CreateEndpoint:    MakeCreateEndpoint(svc),
		DeleteEndpoint:    MakeDeleteEndpoint(svc),
		TotalEndpoint:    MakeTotalEndpoint(svc),
		TopEndpoint:    MakeTopEndpoint(svc),
	}
}

func MakeGetByIDEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
        req:=request.(GetByIDRequest)
		order,err:=svc.GetByID(ctx,req.ID)
		if err != nil {
			return nil, err
		}
		return GetByIDResponse{Order: order},nil
	}
}
func MakeSearchEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
        req:=request.(SearchRequest)
		orders,err:=svc.Search(ctx,req.Search,req.StartDate,req.EndDate)
		if err != nil {
			return nil, err
		}
		return SearchResponse{Orders: orders},nil
	}
}
func MakeCreateEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
        req:=request.(CreateRequest)
		msg,err:=svc.Create(ctx,req.Data)
		if err != nil {
			return nil, err
		}
		return CreateResponse{Message: msg},nil
	}
}
func MakeDeleteEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req:=request.(DeleteRequest)
		msg,err:=svc.Delete(ctx,req.ID)
		if err != nil {
			return nil, err
		}
		return DeleteResponse{Message: msg},nil
	}
}
func MakeTotalEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
        total,err:=svc.Total(ctx)
		if err != nil {
			return nil, err
		}
		return TotalResponse{Total: total},nil
	}
}
func MakeTopEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
        req:=request.(TopRequest)
		products,err:=svc.Top(ctx,req.Count)
		if err != nil {
			return nil, err
		}
		return TopResponse{Products: products},nil
	}
}
