package endpoints

import (
	"CategoryService/internal/database"
	"context"
	"encoding/json"
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

type GetAllRequest struct {

}

type GetAllResponse struct {
	Categories []database.CategoryOut `json:"categories"`
}
type CreateRequest struct {
	Data database.CategoryIn `json:"data"`
}

type CreateResponse struct {
	Message string `json:"message"`
}
type UpdateRequest struct {
	ID uint `json:"id,omitempty"`
	Data database.CategoryIn `json:"data"`
}

type UpdateResponse struct {
	Message string `json:"message"`
}
type DeleteRequest struct {
	ID uint `json:"id,omitempty"`
}

type DeleteResponse struct {
	Message string `json:"message"`
}
type ProductsRequest struct {
	ID uint `json:"id,omitempty"`
}

type ProductsResponse struct {
	Products []database.ProductOut `json:"products"`
}
type GetByIDRequest struct {
	ID uint `json:"id,omitempty"`
}

type GetByIDResponse struct {
	Category database.CategoryOut `json:"category"`
}

type GetByGroupIDRequest struct {
	ID uint `json:"id,omitempty"`
}

type GetByGroupIDResponse struct {
	Categories []database.CategoryOut `json:"categories"`
}


func DecodeGetAllRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request GetAllRequest
	return request,nil
}
func DecodeCreateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request CreateRequest
	err:=json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request,nil
}
func DecodeUpdateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request UpdateRequest
	err:=json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	id,err:=ParseIDFromURL(r)
	if err != nil {
		return request,err
	}
	request.ID=id
	return request,nil
}
func DecodeDeleteRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request DeleteRequest
	err:=json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	id,err:=ParseIDFromURL(r)
	if err != nil {
		return request,err
	}
	request.ID=id
	return request,nil
}
func DecodeProductsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request ProductsRequest
	id,err:=ParseIDFromURL(r)
	if err != nil {
		return request,err
	}
	request.ID=id
	return request,nil
}
func DecodeGetByIDRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request GetByIDRequest
	id,err:=ParseIDFromURL(r)
	if err != nil {
		return request,err
	}
	request.ID=id
	return request,nil
}

func DecodeGetByGroupIDRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request GetByGroupIDRequest
	id,err:=ParseIDFromURL(r)
	if err != nil {
		return request,err
	}
	request.ID=id
	return request,nil
}


func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type","application/json; charset=UTF-8")
	return json.NewEncoder(w).Encode(response)
}