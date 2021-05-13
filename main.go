package main

import (
	"context"
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

	// Initialize MongoDB connection
	db := &database.DB{}
	err := db.Init(context.Background(), database.Config{
		ConnectionString: "mongodb://root:root@localhost:27017",
		DBName:           "dataimpact",
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

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:root@localhost:27017"))
	//if err != nil {
	//    panic(err)
	//}
	//collection := client.Database("dataimpact").Collection("user_auth")
	//
	//ret := db.UserAuthDAO{}
	//
	//result := collection.FindOne(ctx, bson.M{"id": "rooot"})
	//if result.Err() == mongo.ErrNoDocuments {
	//    fmt.Println("HAHAHA")
	//}
	//fmt.Printf("%+v\n", result.Err())
	//
	//if err := result.Decode(&ret); err != nil {
	//    panic(err)
	//}
	//
	//fmt.Printf("%+v", ret)
}
