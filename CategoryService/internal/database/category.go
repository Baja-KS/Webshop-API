package database

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"os"
	"reflect"
	"strconv"
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
	GroupID uint `json:"GroupId"`
}

type CategoryOut struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
	GroupID uint `json:"GroupId"`
	Deletable bool `json:"deletable"`
}

//func GetCategories(GroupID uint,categoryServiceURL string) ([]CategoryOut,error) {
//
//	var categories []CategoryOut
//	var response CategoryServiceResponse
//	res,err:=http.Get(categoryServiceURL+"/GetByGroupID/"+ strconv.Itoa(int(GroupID)))
//	if err != nil {
//		return categories,err
//	}
//	err=json.NewDecoder(res.Body).Decode(&response)
//	if err != nil {
//		return categories,err
//	}
//	return response.Categories,nil
//}

type ProductServiceResponse struct {
	Products []ProductOut `json:"products"`
}

func (c *Category) GetProducts(productServiceURL string) ([]ProductOut, error) {
	var products []ProductOut
	var response ProductServiceResponse
	res,err:=http.Get(productServiceURL+"/Search?CategoryId="+strconv.Itoa(int(c.ID)))
	if err != nil {
		return products,err
	}
	err=json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return products,err
	}
	return response.Products,nil
}

func (c *Category) IsDeletable() bool {
	products,err:=c.GetProducts(os.Getenv("PRODUCT_SERVICE"))
	if err != nil {
		return false
	}
	return len(products)==0
}

func (c *Category) Out() CategoryOut {
	return CategoryOut{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		GroupID:     c.GroupID,
		Deletable: c.IsDeletable(),
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
