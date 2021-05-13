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
	mongoConn   = os.Getenv("MONGO_CONNECTION_STRING")
	mongoDBName = os.Getenv("MONGO_DB_NAME")
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
		SigningKey:    []byte("dataimpact"),
	}

	// Initialize HTTP handler
	h := handler.Handler{}
	h.Init(db, fs, t)

	e.POST("/login", h.Login)
	e.POST("/create", h.CreateUser)
	e.GET("/user/:id", h.GetUser)

	e.Logger.Fatal(e.Start(":8080"))
}
