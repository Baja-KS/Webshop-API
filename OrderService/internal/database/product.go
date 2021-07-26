package database


type ProductOut struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
	Img string `json:"img,omitempty"`
	Price float32 `json:"price"`
	Discount int `json:"discount"`
	CategoryID uint `json:"categoryID"`
}