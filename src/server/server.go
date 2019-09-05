package server

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"time"
)

var (
	serverDoneSignal = make(chan struct{})
)

func Run(stopRequest chan struct{}) {
	defer func() {
		close(serverDoneSignal)
	}()

	keys, err := ReadKeys()
	if err != nil {
		log.Println(err.Error())
		return
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.POST("/login", Login(keys))

	g := e.Group("/api")
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    keys.PublicKey,
		SigningMethod: SignMethod.Name,
		Claims:        &CustomClaims{},
	}))
	g.POST("/refresh", Refresh(keys))
	g.POST("/query", Query())

	go func() {
		// 実際にはここでStartTSLで開始する
		err := e.Start(":8888")
		e.Logger.Fatal(err)
	}()

	<-stopRequest
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err = e.Shutdown(ctx)
	if err != nil {
		e.Logger.Fatal(err)
	}

	return
}

func Wait() {
	<-serverDoneSignal
}

func GetCustomClaims(c echo.Context) (*CustomClaims, error) {
	token := c.Get("user").(*jwt.Token)
	CustomClaims := token.Claims.(*CustomClaims)
	return CustomClaims, nil
}
