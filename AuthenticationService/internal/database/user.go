package database

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID uint `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	Username string `gorm:"not null;unique" json:"username"`
	Fullname string `gorm:"not null" json:"fullname"`
	Email string `gorm:"not null;unique" json:"email"`
	Password string `gorm:"not null"  json:"password"`
	IsAdmin bool `gorm:"default:false" json:"isAdmin"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
}

func (u *User) Out() (UserOut,error) {
	if u.Username=="" {
		var emptyUser UserOut
		return emptyUser,errors.New("Non existent user")
	}
	userOut:=UserOut{
		ID:        u.ID,
		Username:  u.Username,
		Fullname:  u.Fullname,
		Email:     u.Email,
		IsAdmin:   u.IsAdmin,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	return userOut,nil
}

type UserIn struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email string `json:"email"`
	Password string `json:"password"`
	IsAdmin bool `json:"isAdmin"`
}

type UserOut struct {
	ID uint `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email string `json:"email"`
	IsAdmin bool `json:"isAdmin"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}