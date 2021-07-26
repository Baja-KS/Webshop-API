package service

import (
	"OrderService/internal/database"
	"context"
	"gorm.io/gorm"
	"time"
)

//OrderService should implement the Service interface


type OrderService struct {
	DB *gorm.DB
}

type Service interface {
	GetByID(ctx context.Context,id uint) (database.OrderOut,error)
	Search(ctx context.Context,search string,startDate time.Time,endDate time.Time) ([]database.OrderOut,error)
	Create(ctx context.Context,data database.OrderIn) (string,error)
	Delete(ctx context.Context,id uint) (string,error)
	Total(ctx context.Context) (float32,error)
	Top(ctx context.Context,count uint) ([]database.ProductOut,error)
}