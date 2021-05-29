package middlewares

import (
	"github.com/Baja-KS/Webshop-API/AuthenticationService/internal/database"
	"github.com/Baja-KS/Webshop-API/AuthenticationService/internal/service"
	"context"
	"github.com/go-kit/kit/log"
	"time"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   service.Service
}

func (l *LoggingMiddleware) Login(ctx context.Context, username string, password string) (user database.User,token string,err error) {
	defer func(begin time.Time) {
		err := l.Logger.Log("method", "login", "user", user.Fullname, "token", token, "err", err, "took", time.Since(begin))
		if err != nil {
			return
		}
	}(time.Now())
	user,token,err=l.Next.Login(ctx,username,password)
	return
}

func (l *LoggingMiddleware) Register(ctx context.Context, user database.UserIn) (msg string, err error) {
	defer func(begin time.Time){
		err := l.Logger.Log("method", "register", "msg",msg,"err", err, "took", time.Since(begin))
		if err != nil {
			return
		}
	}(time.Now())
	msg,err=l.Next.Register(ctx,user)
	return
}

func (l *LoggingMiddleware) GetAll(ctx context.Context) (users []database.User,err error) {
	defer func(begin time.Time){
		err := l.Logger.Log("method", "getall", "err", err, "took", time.Since(begin))
		if err != nil {
			return
		}
	}(time.Now())
	users,err=l.Next.GetAll(ctx)
	return
}

func (l *LoggingMiddleware) AuthUser(ctx context.Context, tokenString string) (user database.User, err error) {
	defer func(begin time.Time){
		err := l.Logger.Log("method", "authuser","token",tokenString,"user",user.Fullname,"err", err, "took", time.Since(begin))
		if err != nil {
			return
		}
	}(time.Now())
	user,err=l.Next.AuthUser(ctx, tokenString)
	return
}