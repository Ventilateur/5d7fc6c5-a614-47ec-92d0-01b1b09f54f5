package main

import (
	"context"
	"github.com/Ventilateur/dataimpact-test/config"
	"github.com/Ventilateur/dataimpact-test/database"
	"github.com/Ventilateur/dataimpact-test/filesystem"
	"github.com/Ventilateur/dataimpact-test/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"time"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Read config
	conf, err := config.GetConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Initialize MongoDB connection
	db := &database.DB{}
	if err = db.Init(context.Background(), database.Config{
		ConnectionString: conf.Mongo.ConnectionString,
		DBName:           conf.Mongo.DatabaseName,
	}); err != nil {
		e.Logger.Fatal(err)
	}

	// Initialize file storage
	fs := &filesystem.FileStorage{}
	if err = fs.Init(conf.UserDataDir); err != nil {
		e.Logger.Fatal(err)
	}

	// Initialize JWT tokenizer
	t := &handler.Tokenizer{
		TokenDuration: time.Minute * time.Duration(conf.JWT.DurationMin),
		SigningKey:    []byte(conf.JWT.SigningKey),
	}

	// Initialize HTTP handler
	h := handler.Handler{}
	h.Init(db, fs, t)

	// JWT authorization
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(conf.JWT.SigningKey),
		TokenLookup: "header:" + echo.HeaderAuthorization,
		Skipper: func(c echo.Context) bool {
			return "/login" == c.Request().URL.Path
		},
	}))

	e.POST("/login", h.Login)
	e.POST("/create", h.CreateUser)
	e.GET("/user/:id", h.GetUser)
	e.GET("/users/list", h.ListUsers)
	e.DELETE("/delete/user/:id", h.DeleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}
