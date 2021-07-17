package service

import (
	"CategoryService/internal/database"
	"context"
	"errors"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strconv"
)

//CategoryService should implement the Service interface


type CategoryService struct {
	DB *gorm.DB
}

func ValidateGroup(groupServiceURL string, id uint) bool {
	_,err:=http.Get(groupServiceURL+"/GetByID/"+ strconv.Itoa(int(id)))
	if err != nil {
		return false
	}
	return true
}

type Service interface {
	GetAll(ctx context.Context) ([]database.CategoryOut,error)
	Create(ctx context.Context,data database.CategoryIn) (string,error)
	Update(ctx context.Context,id uint,data database.CategoryIn) (string,error)
	Delete(ctx context.Context,id uint) (string,error)
	Products(ctx context.Context,id uint) ([]database.ProductOut,error)
	GetByID(ctx context.Context,id uint) (database.CategoryOut,error)
	GetByGroupID(ctx context.Context,id uint) ([]database.CategoryOut,error)
}

func (c *CategoryService) GetAll(ctx context.Context) ([]database.CategoryOut, error) {
	var categories []database.Category
	result:=c.DB.Find(&categories)
	if result.Error != nil {
		return database.CategoryArrayOut(categories),result.Error
	}
	out:=database.CategoryArrayOut(categories)
	return out,nil
}

func (c *CategoryService) Create(ctx context.Context, data database.CategoryIn) (string, error) {
	category:=data.In()
	if !ValidateGroup(os.Getenv("GROUP_SERVICE"),category.GroupID) {
		return "Non existent group",errors.New("group with that ID doesnt exist")
	}
	result:=c.DB.Create(&category)
	if result.Error != nil {
		return "Error", result.Error
	}
	return "Successfully created", nil
}

func (c *CategoryService) Update(ctx context.Context, id uint, data database.CategoryIn) (string, error) {
	var category database.Category
	notFound:=c.DB.Where("id = ?",id).First(&category).Error
	if notFound != nil {
		return "That category doesn't exist", notFound
	}
	if !ValidateGroup(os.Getenv("GROUP_SERVICE"),data.GroupID) {
		return "Non existent group",errors.New("group with that ID doesnt exist")
	}
	category=category.Update(data)
	err:=c.DB.Save(&category).Error
	if err != nil {
		return "Error updating category", err
	}

	return "Category updated successfully", nil
}

func (c *CategoryService) Delete(ctx context.Context, id uint) (string, error) {
	var category database.Category
	notFound:=c.DB.Where("id = ?",id).First(&category).Error
	if notFound != nil {
		return "That category doesn't exist", notFound
	}
	err:=c.DB.Delete(&database.Category{},id).Error
	if err != nil {
		return "Error deleting category", err
	}
	return "Category successfully deleted", nil
}

func (c *CategoryService) Products(ctx context.Context, id uint) ([]database.ProductOut, error) {
	panic("implement me")
}

func (c *CategoryService) GetByID(ctx context.Context, id uint) (database.CategoryOut, error) {
	var category database.Category
	c.DB.Where("id = ?",id).First(&category)
	return category.Out(),nil
}

func (c *CategoryService) GetByGroupID(ctx context.Context, id uint) ([]database.CategoryOut, error) {
	var categories []database.Category
	result:=c.DB.Where("group_id = ?",id).Find(&categories)
	if result.Error != nil {
		return database.CategoryArrayOut(categories),result.Error
	}
	out:=database.CategoryArrayOut(categories)
	return out,nil
}
