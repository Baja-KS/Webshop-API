package service

import (
	"GroupService/internal/database"
	"context"
	"gorm.io/gorm"
)

//GroupService should implement the Service interface



type GroupService struct {
	DB *gorm.DB
}

func (g *GroupService) GetAll(context.Context) ([]database.GroupOut, error) {
	var groups []database.Group
	result:=g.DB.Find(&groups)
	if result.Error != nil {
		return database.GroupArrayOut(groups), result.Error
	}
	out:=database.GroupArrayOut(groups)
	return out,nil
}

func (g *GroupService) Create(ctx context.Context, data database.GroupIn) (string, error) {
	group:=data.In()
	result:=g.DB.Create(&group)
	if result.Error != nil {
		return "Error", result.Error
	}
	return "Successfully created", nil
}

func (g *GroupService) Update(ctx context.Context, id uint, data database.GroupIn) (string, error) {
	var group database.Group
	notFound:=g.DB.Where("id = ?",id).First(&group).Error
	if notFound != nil {
		return "That group doesn't exist", notFound
	}
	if data.Name!=""  {
		group.Name=data.Name
	}
	group.Description=data.Description
	err:=g.DB.Save(&group).Error
	if err != nil {
		return "Error updating group", err
	}

	return "Group updated successfully", nil
}

func (g *GroupService) Delete(ctx context.Context, id uint) (string, error) {
	var group database.Group
	notFound:=g.DB.Where("id = ?",id).First(&group).Error
	if notFound != nil {
		return "That group doesn't exist", notFound
	}
	err:=g.DB.Delete(&database.Group{},id).Error
	if err != nil {
		return "Error deleting group", err
	}
	return "Group successfully deleted", nil
}

func (g *GroupService) Categories(ctx context.Context, id uint) ([]database.CategoryOut, error) {
	panic("implement me")
}

func (g *GroupService) GetByID(ctx context.Context, id uint) (database.GroupOut, error) {
	var group database.Group
	g.DB.Where("id = ?",id).First(&group)
	return group.Out(),nil
}

type Service interface {
	GetAll(ctx context.Context) ([]database.GroupOut, error)
	Create(ctx context.Context, data database.GroupIn) (string, error)
	Update(ctx context.Context, id uint, data database.GroupIn) (string, error)
	Delete(ctx context.Context, id uint) (string, error)
	Categories(ctx context.Context, id uint) ([]database.CategoryOut, error)
	GetByID(ctx context.Context, id uint) (database.GroupOut, error)

}

