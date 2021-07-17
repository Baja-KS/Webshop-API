package database

import (
	"gorm.io/gorm"
	"reflect"
)


type Category struct {
	gorm.Model
	ID uint `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	Name string `gorm:"not null;unique" json:"name"`
	Description string `gorm:"" json:"description,omitempty"`
	GroupID uint `gorm:"not null" json:"groupID"`
}

type CategoryIn struct {
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
	GroupID uint `json:"groupID"`
}

type CategoryOut struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
	GroupID uint `json:"groupID"`
}


func (c *Category) Out() CategoryOut {
	return CategoryOut{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		GroupID:     c.GroupID,
	}
}

func (c *Category) Update(data CategoryIn) Category {
	updated:=*c
	forUpdate:=reflect.ValueOf(data)
	for i:=0;i<forUpdate.NumField();i++ {
		field:=forUpdate.Type().Field(i).Name
		value:=forUpdate.Field(i)
		v := reflect.ValueOf(&updated).Elem().FieldByName(field)
		if v.IsValid() {
			v.Set(value)
		}

	}
	return updated
}

func CategoryArrayOut(categoryModels []Category) []CategoryOut {
	outArr:=make([]CategoryOut,len(categoryModels))
	for i,category := range categoryModels {
		outArr[i]=category.Out()
	}
	return outArr
}

func (i *CategoryIn) In() Category {
	return Category{
		Name:        i.Name,
		Description: i.Description,
		GroupID:     i.GroupID,
	}
}
