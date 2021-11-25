package endpoints

import (
	"GroupService/internal/database"
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
	Groups []database.GroupOutWithCategories `json:"groups"`
}
type CreateRequest struct {
	Input database.GroupIn `json:"input"`
}

type CreateResponse struct {
 	Message string `json:"message"`
}
type UpdateRequest struct {
	ID uint `json:"id,omitempty"`
	Data database.GroupIn `json:"data"`
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
type CategoriesRequest struct {
	ID uint `json:"id,omitempty"`
}

type CategoriesResponse struct {
	Categories []database.CategoryOut `json:"categories"`
}
type GetByIDRequest struct {
	ID uint `json:"id,omitempty"`
}

type GetByIDResponse struct {
	Group database.GroupOut `json:"group"`
}


func DecodeGetAllRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request GetAllRequest
	return request,nil
}
func DecodeCreateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request CreateRequest
	err:=json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return request,err
	}
	return request,nil
}
func DecodeUpdateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request UpdateRequest
	err:=json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		return request,err
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

	id,err:=ParseIDFromURL(r)
	if err != nil {
		return request,err
	}
	request.ID=id
	return request,nil
}
func DecodeCategoriesRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request CategoriesRequest


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


func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type","application/json; charset=UTF-8")
	return json.NewEncoder(w).Encode(response)
}