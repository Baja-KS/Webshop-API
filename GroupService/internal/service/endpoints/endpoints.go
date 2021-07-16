package endpoints

import (
	"GroupService/internal/service"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type EndpointSet struct {
	GetAllEndpoint endpoint.Endpoint
	CreateEndpoint endpoint.Endpoint
	UpdateEndpoint endpoint.Endpoint
	DeleteEndpoint endpoint.Endpoint
	CategoriesEndpoint endpoint.Endpoint
	GetByIDEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc service.Service) EndpointSet {
	return EndpointSet{
		GetAllEndpoint:    MakeGetAllEndpoint(svc),
		CreateEndpoint:    MakeCreateEndpoint(svc),
		UpdateEndpoint:    MakeUpdateEndpoint(svc),
		DeleteEndpoint:    MakeDeleteEndpoint(svc),
		CategoriesEndpoint:    MakeCategoriesEndpoint(svc),
		GetByIDEndpoint:    MakeGetByIDEndpoint(svc),
	}
}

func MakeGetAllEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		groups,err:=svc.GetAll(ctx)
		if err != nil {
			return nil, err
		}
		return GetAllResponse{Groups: groups},nil
	}
}
func MakeCreateEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
        req:=request.(CreateRequest)
		msg,err:=svc.Create(ctx, req.Input)
		if err != nil {
			return nil, err
		}
		return CreateResponse{Message: msg},nil
	}
}
func MakeUpdateEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
        req:=request.(UpdateRequest)
		msg,err:=svc.Update(ctx, req.ID, req.Data)
		if err != nil {
			return nil, err
		}
		return UpdateResponse{Message: msg},nil
	}
}
func MakeDeleteEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
        req:=request.(DeleteRequest)
		msg,err:=svc.Delete(ctx, req.ID)
		if err != nil {
			return nil, err
		}
		return DeleteResponse{Message: msg},nil
	}
}
func MakeCategoriesEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
        //req:=request.(CategoriesRequest)

		return CategoriesResponse{},nil
	}
}
func MakeGetByIDEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
        req:=request.(GetByIDRequest)
		group,err:=svc.GetByID(ctx, req.ID)
		if err != nil {
			return nil, err
		}
		return GetByIDResponse{Group: group},nil
	}
}
