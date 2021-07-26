package database

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	ID uint `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	Count uint `gorm:"default:1" json:"count"`
	OrderID uint `gorm:"not null" json:"orderID"`
	ProductID uint `gorm:"not null" json:"productID"`
}

type OrderItemIn struct {
	Count uint `json:"count"`
	ProductID uint `json:"productID"`
}

type OrderItemOut struct {
	ID uint `json:"id,omitempty"`
	Count uint `json:"count"`
	OrderID uint `json:"orderID"`
	ProductID uint `json:"productID"`
}

func (i *OrderItem) Out() OrderItemOut {
	return OrderItemOut{
		ID:        i.ID,
		Count:     i.Count,
		OrderID:   i.OrderID,
		ProductID: i.ProductID,
	}
}

func (i *OrderItemIn) In() OrderItem {
	return OrderItem{
		Count:     i.Count,
		ProductID: i.ProductID,
	}
}

func ItemArrayOut(models []OrderItem) []OrderItemOut {
	outArr:=make([]OrderItemOut,len(models))
	for i,item := range models {
		outArr[i]=item.Out()
	}
	return outArr
}

func ItemArrayIn(models []OrderItemIn) []OrderItem {
	inArr:=make([]OrderItem,len(models))
	for i, item := range models {
		inArr[i]=item.In()
	}
	return inArr
}

