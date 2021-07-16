package service

import (
	"context"
	"errors"
	"github.com/Baja-KS/Webshop-API/AuthenticationService/internal/database"
	stdjwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
	"strconv"
)
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type Service interface {
	Login(ctx context.Context,username string,password string) (database.User,string,error)
	Register(ctx context.Context, user database.UserIn) (string, error)
	GetAll(ctx context.Context) ([]database.User,error)
	AuthUser(ctx context.Context, tokenString string) (database.User, error)
}


func (a *AuthenticationService) Login(ctx context.Context, username string, password string) (database.User, string, error) {
	var userDB database.User
	err:=a.DB.Where("username = ?",username).First(&userDB).Error
	if err != nil {
		return database.User{}, "", err
	}
	if !CheckPasswordHash(password, userDB.Password) {
		return database.User{}, "Wrong password", nil
	}
	//atClaims:=stdjwt.MapClaims{}
	//atClaims["authorized"]=true
	//atClaims["id"]=userDB.ID
	at:=stdjwt.NewWithClaims(stdjwt.SigningMethodHS256,stdjwt.StandardClaims{
		Id: strconv.Itoa(int(userDB.ID)),
	})
	token,err:=at.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return database.User{}, "", err
	}
	return userDB,token,nil
}

func (a *AuthenticationService) Register(ctx context.Context, user database.UserIn) (string, error) {
	hash,e:=HashPassword(user.Password)
	if e != nil {
		return "Error",e
	}
	dbUser:=database.User{
		Username:  user.Username,
		Fullname:  user.Fullname,
		Email:     user.Email,
		Password:  hash,
		IsAdmin:   false,
	}
	result:=a.DB.Create(&dbUser)
	if result.Error != nil {
		return "Error",result.Error
	}
	return "Register successful",nil
}

func (a *AuthenticationService) GetAll(ctx context.Context) ([]database.User, error) {

	var users []database.User
	result:=a.DB.Find(&users)

	if result.Error != nil {
		return users,result.Error
	}
	return users,nil
}

func (a *AuthenticationService) AuthUser(ctx context.Context, tokenString string) (database.User, error) {
	//claims:=stdjwt.MapClaims{}
	parsed,err:=stdjwt.ParseWithClaims(tokenString,&stdjwt.StandardClaims{}, func(token *stdjwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")),nil
	})
	if err != nil {
		return database.User{}, err
	}
	if parsed.Valid {
		castedClaims:=parsed.Claims.(*stdjwt.StandardClaims)
		userId:=castedClaims.Id
		id,err:=strconv.ParseUint(userId,10,64)
		if err != nil {
			return database.User{}, err
		}
		var userDB database.User
		notFound:=a.DB.Where("id = ?",id).First(&userDB).Error
		if notFound != nil {
			return database.User{}, notFound
		}
		return userDB,nil
	}
	return database.User{}, errors.New("Not found")
}

type AuthenticationService struct {
	DB *gorm.DB
}

