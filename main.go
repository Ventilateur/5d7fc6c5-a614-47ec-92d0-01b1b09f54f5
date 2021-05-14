package main

import (
	"context"
	"github.com/Ventilateur/dataimpact-test/database"
	"github.com/Ventilateur/dataimpact-test/filesystem"
	"github.com/Ventilateur/dataimpact-test/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
	"time"
)

var (
	mongoConn     = os.Getenv("MONGO_CONNECTION_STRING")
	mongoDBName   = os.Getenv("MONGO_DB_NAME")
	jwtSigningKey = []byte(os.Getenv("JWT_SIGNING_KEY"))
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize MongoDB connection
	db := &database.DB{}
	err := db.Init(context.Background(), database.Config{
		ConnectionString: mongoConn,
		DBName:           mongoDBName,
	})
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Initialize file storage
	fs := &filesystem.FileStorage{}
	if err = fs.Init("userdata"); err != nil {
		e.Logger.Fatal(err)
	}

	// Initialize JWT tokenizer
	t := &handler.Tokenizer{
		TokenDuration: time.Minute * 10,
		SigningKey:    jwtSigningKey,
	}

	// Initialize HTTP handler
	h := handler.Handler{}
	h.Init(db, fs, t)

	// JWT authorization
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  jwtSigningKey,
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
