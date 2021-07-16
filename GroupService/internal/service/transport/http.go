package transport

import (
	"GroupService/internal/service/endpoints"
	"context"
	"errors"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strings"
)

func GetAuthToken(r *http.Request) (string,error) {
	authHeader:=r.Header["Authorization"]
	if len(authHeader)==0 {
		return "", errors.New("no auth header")
	}
	authHeaderParts:=strings.Split(authHeader[0]," ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "",errors.New("invalid auth header")
	}
	return authHeaderParts[1],nil
}

func AuthExtractor(ctx context.Context, r *http.Request) context.Context {
	token,err:=GetAuthToken(r)
	if err != nil {
		return context.WithValue(ctx,"auth","")
	}
	return context.WithValue(ctx,"auth",token)
}

func NewHTTPHandler(ep endpoints.EndpointSet) http.Handler {
	router:=mux.NewRouter()

	GetAllHandler:=httptransport.NewServer(ep.GetAllEndpoint,endpoints.DecodeGetAllRequest,endpoints.EncodeResponse,httptransport.ServerBefore(AuthExtractor))
	CreateHandler:=httptransport.NewServer(ep.CreateEndpoint,endpoints.DecodeCreateRequest,endpoints.EncodeResponse,httptransport.ServerBefore(AuthExtractor))
	UpdateHandler:=httptransport.NewServer(ep.UpdateEndpoint,endpoints.DecodeUpdateRequest,endpoints.EncodeResponse,httptransport.ServerBefore(AuthExtractor))
	DeleteHandler:=httptransport.NewServer(ep.DeleteEndpoint,endpoints.DecodeDeleteRequest,endpoints.EncodeResponse,httptransport.ServerBefore(AuthExtractor))
	CategoriesHandler:=httptransport.NewServer(ep.CategoriesEndpoint,endpoints.DecodeCategoriesRequest,endpoints.EncodeResponse,httptransport.ServerBefore(AuthExtractor))
	GetByIDHandler:=httptransport.NewServer(ep.GetByIDEndpoint,endpoints.DecodeGetByIDRequest,endpoints.EncodeResponse,httptransport.ServerBefore(AuthExtractor))


	router.Handle("/GetAll",GetAllHandler).Methods(http.MethodGet)
	router.Handle("/Create",CreateHandler).Methods(http.MethodPost)
	router.Handle("/Update/{id}",UpdateHandler).Methods(http.MethodPut)
	router.Handle("/Delete/{id}",DeleteHandler).Methods(http.MethodDelete)
	router.Handle("/Categories",CategoriesHandler).Methods(http.MethodGet)
	router.Handle("/GetByID/{id}",GetByIDHandler).Methods(http.MethodGet)
	router.Handle("/metrics",promhttp.Handler()).Methods(http.MethodGet)


	return router
}