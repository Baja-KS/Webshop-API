package middlewares

import (
	"GroupService/internal/database"
	"GroupService/internal/service"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
)


func CheckAuth(ctx context.Context,authServiceURL string) bool {
	client:=&http.Client{}
	req,err:=http.NewRequest("GET",authServiceURL+"/User",nil)
	if err != nil {
		return false
	}
	token:=ctx.Value("auth").(string)
	authHeader:=fmt.Sprintf("Bearer %s",token)
	req.Header.Add("Authorization",authHeader)
	res, err := client.Do(req)
	if err != nil || res.StatusCode!=200 {
		return false
	}

	return true
}


type AuthenticationMiddleware struct {
	Next service.Service
}

func (a *AuthenticationMiddleware) GetAll(ctx context.Context) ([]database.GroupOutWithCategories, error) {
	return a.Next.GetAll(ctx)
}

func (a *AuthenticationMiddleware) Create(ctx context.Context, data database.GroupIn) (string, error) {
	if !CheckAuth(ctx,os.Getenv("AUTH_SERVICE"))  {
		return "Unauthorized",errors.New("unauthorized")
	}
	return a.Next.Create(ctx,data)
}

func (a *AuthenticationMiddleware) Update(ctx context.Context, id uint, data database.GroupIn) (string, error) {
	if !CheckAuth(ctx,os.Getenv("AUTH_SERVICE"))  {
		return "Unauthorized",errors.New("unauthorized")
	}
	return a.Next.Update(ctx,id,data)
}

func (a *AuthenticationMiddleware) Delete(ctx context.Context, id uint) (string, error) {
	if !CheckAuth(ctx,os.Getenv("AUTH_SERVICE"))  {
		return "Unauthorized",errors.New("unauthorized")
	}
	return a.Next.Delete(ctx,id)
}

func (a *AuthenticationMiddleware) Categories(ctx context.Context, id uint) ([]database.CategoryOut, error) {
	return a.Next.Categories(ctx,id)
}

func (a *AuthenticationMiddleware) GetByID(ctx context.Context, id uint) (database.GroupOut, error) {
	return a.Next.GetByID(ctx,id)
}

