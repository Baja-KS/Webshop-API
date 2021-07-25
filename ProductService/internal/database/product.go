package database

import (
	"gorm.io/gorm"
	"io"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"strconv"
)

func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

type Product struct {
	gorm.Model
	ID uint `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	Name string `gorm:"not null;unique" json:"name"`
	Description string `gorm:"" json:"description,omitempty"`
	Img string `gorm:"" json:"img,omitempty"`
	Price float32 `gorm:"not null" json:"price"`
	Discount int `gorm:"" json:"discount"`
	CategoryID uint `gorm:"not null" json:"categoryID"`
}

type ProductIn struct {
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
	Img string `json:"img,omitempty"`
	Price float32 `json:"price"`
	Discount int `json:"discount"`
	CategoryID uint `json:"categoryID"`
}

type ProductOut struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
	Img string `json:"img,omitempty"`
	Price float32 `json:"price"`
	Discount int `json:"discount"`
	CategoryID uint `json:"categoryID"`
}

func (p *Product) Out() ProductOut {
	return ProductOut{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Img:         p.Img,
		Price:       p.Price,
		Discount:    p.Discount,
		CategoryID:  p.CategoryID,
	}
}

func (i *ProductIn) In() Product {
	return Product{
		Name:        i.Name,
		Description: i.Description,
		Img:         i.Img,
		Price:       i.Price,
		Discount:    i.Discount,
		CategoryID:  i.CategoryID,
	}
}

func ProductArrayOut(models []Product) []ProductOut {
	outArr:=make([]ProductOut,len(models))
	for i,item := range models {
		outArr[i]=item.Out()
	}
	return outArr
}

func (p *Product) Update(data ProductIn) Product {
	updated:=*p
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

func DecodeMultipartRequest(r *http.Request) (ProductIn,error) {
	err:=r.ParseMultipartForm(32<<20)
	if err != nil {
		return ProductIn{}, err
	}
	//text fields
	price,err:=strconv.ParseFloat(r.FormValue("price"),32)
	if err != nil {
		return ProductIn{}, err
	}
	discount,err:=strconv.Atoi(r.FormValue("discount"))
	if err != nil {
		return ProductIn{}, err
	}
	category,err:=strconv.ParseUint(r.FormValue("CategoryID"),10,32)
	if err != nil {
		return ProductIn{}, err
	}
	data:=ProductIn{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Price:       float32(price),
		Discount:    discount,
		CategoryID:  uint(category),
	}
	file, _,err:=r.FormFile("img")
	if err != nil {
		return data, nil
	}
	err = file.Close()
	if err != nil {
		return data, err
	}
	imgName:=RandomString(16)
	f,err:=os.OpenFile("../../resources/img/"+imgName,os.O_WRONLY|os.O_CREATE,0777)
	if err != nil {
		return data, err
	}
	_, err = io.Copy(f,file)
	err=f.Close()
	if err != nil {
		return data, err
	}
	data.Img=imgName
	return data,nil
}
