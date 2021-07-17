package endpoints

import (
	"CategoryService/internal/service"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type EndpointSet struct {
	GetAllEndpoint endpoint.Endpoint
	CreateEndpoint endpoint.Endpoint
	UpdateEndpoint endpoint.Endpoint
	DeleteEndpoint endpoint.Endpoint
	ProductsEndpoint endpoint.Endpoint
	GetByIDEndpoint endpoint.Endpoint
	GetByGroupIDEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc service.Service) EndpointSet {
	return EndpointSet{
		GetAllEndpoint:    MakeGetAllEndpoint(svc),
		CreateEndpoint:    MakeCreateEndpoint(svc),
		UpdateEndpoint:    MakeUpdateEndpoint(svc),
		DeleteEndpoint:    MakeDeleteEndpoint(svc),
		ProductsEndpoint:    MakeProductsEndpoint(svc),
		GetByIDEndpoint:    MakeGetByIDEndpoint(svc),
		GetByGroupIDEndpoint: MakeGetByGroupIDEndpoint(svc),
	}
}

func MakeGetAllEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		categories,err:=svc.GetAll(ctx)
		if err != nil {
			return nil, err
		}
		return GetAllResponse{Categories: categories},nil
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
func MakeUpdateEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
        req:=request.(UpdateRequest)
		msg,err:=svc.Update(ctx,req.ID,req.Data)
		if err != nil {
			return nil, err
		}
		return UpdateResponse{Message: msg},nil
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
func MakeProductsEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
        req:=request.(ProductsRequest)
		products,err:=svc.Products(ctx,req.ID)
		if err != nil {
			return nil, err
		}
		return ProductsResponse{Products: products},nil
	}
}
func MakeGetByIDEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
        req:=request.(GetByIDRequest)
		category,err:=svc.GetByID(ctx,req.ID)
		if err != nil {
			return nil, err
		}
		return GetByIDResponse{Category: category},nil
	}
}

func MakeGetByGroupIDEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req:=request.(GetByGroupIDRequest)
		categories,err:=svc.GetByGroupID(ctx,req.ID)
		if err != nil {
			return nil, err
		}
		return GetByGroupIDResponse{Categories: categories},nil
	}
}
