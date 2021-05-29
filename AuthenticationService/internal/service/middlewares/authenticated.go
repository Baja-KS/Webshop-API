package middlewares

import (
	stdjwt "github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"os"
)



func Authenticated() endpoint.Middleware {
	kf:=func(token *stdjwt.Token) (interface{}, error) { return []byte(os.Getenv("JWT_KEY")), nil }
	return jwt.NewParser(kf,stdjwt.SigningMethodHS256,jwt.StandardClaimsFactory)
}
