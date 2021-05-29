package transport

import (
	"github.com/Baja-KS/Webshop-API/AuthenticationService/internal/service/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHTTPHandler(ep endpoints.EndpointSet) http.Handler {
	router:=mux.NewRouter()

	loginHandler:=httptransport.NewServer(ep.LoginEndpoint,endpoints.DecodeLoginRequest,endpoints.EncodeResponse)
	registerHandler:=httptransport.NewServer(ep.RegisterEndpoint,endpoints.DecodeRegisterRequest,endpoints.EncodeResponse)
	getAllHandler:=httptransport.NewServer(ep.GetAllEndpoint,endpoints.DecodeGetAllRequest,endpoints.EncodeResponse)
	authUserHandler:=httptransport.NewServer(ep.AuthUserEndpoint,endpoints.DecodeAuthUserRequest,endpoints.EncodeResponse)
	router.Handle("/auth/login",loginHandler).Methods(http.MethodPost)
	router.Handle("/auth/register",registerHandler).Methods(http.MethodPost)
	router.Handle("/user/getAll",getAllHandler).Methods(http.MethodGet)
	router.Handle("/auth/user",authUserHandler).Methods(http.MethodGet)

	return router
}