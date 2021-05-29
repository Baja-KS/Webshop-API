package endpoints

import (
	"github.com/Baja-KS/Webshop-API/AuthenticationService/internal/database"
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User database.UserOut `json:"user,omitempty"`
	Token string `json:"token,omitempty"`
	Message string `json:"message"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Password string `json:"password"`
	Email string `json:"email"`
	IsAdmin bool `json:"isAdmin"`
}

type RegisterResponse struct {
	//ID uint `json:"id"`
	Username string `json:"username"`
	Message string `json:"message"`
}

type GetAllRequest struct {
	
}

type GetAllResponse struct {
	Users []database.UserOut `json:"users"`
}

type AuthUserRequest struct {
	Token string
}

type AuthUserResponse struct {
	User database.UserOut `json:"user,omitempty"`
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type","application/json; charset=UTF-8")
	return json.NewEncoder(w).Encode(response)
}

func DecodeLoginRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request LoginRequest
	err:=json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request,nil
}

func DecodeRegisterRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request RegisterRequest
	err:=json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request,nil
}

func DecodeGetAllRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request GetAllRequest
	return request,nil
}
func DecodeAuthUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request AuthUserRequest
	authHeader:=r.Header["Authorization"]
	authHeaderParts:=strings.Split(authHeader[0]," ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return request,nil
	}
	request.Token=authHeaderParts[1]
	return request,nil
}
