package transport

import (
	"OrderService/internal/service/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHTTPHandler(ep endpoints.EndpointSet) http.Handler {
	router:=mux.NewRouter()

	GetByIDHandler:=httptransport.NewServer(ep.GetByIDEndpoint,endpoints.DecodeGetByIDRequest,endpoints.EncodeResponse) 
	SearchHandler:=httptransport.NewServer(ep.SearchEndpoint,endpoints.DecodeSearchRequest,endpoints.EncodeResponse) 
	CreateHandler:=httptransport.NewServer(ep.CreateEndpoint,endpoints.DecodeCreateRequest,endpoints.EncodeResponse) 
	DeleteHandler:=httptransport.NewServer(ep.DeleteEndpoint,endpoints.DecodeDeleteRequest,endpoints.EncodeResponse) 
	TotalHandler:=httptransport.NewServer(ep.TotalEndpoint,endpoints.DecodeTotalRequest,endpoints.EncodeResponse) 
	TopHandler:=httptransport.NewServer(ep.TopEndpoint,endpoints.DecodeTopRequest,endpoints.EncodeResponse) 


	router.Handle("/GetByID",GetByIDHandler).Methods(http.MethodGet)
	router.Handle("/Search",SearchHandler).Methods(http.MethodGet)
	router.Handle("/Create",CreateHandler).Methods(http.MethodGet)
	router.Handle("/Delete",DeleteHandler).Methods(http.MethodGet)
	router.Handle("/Total",TotalHandler).Methods(http.MethodGet)
	router.Handle("/Top",TopHandler).Methods(http.MethodGet)


	return router
}