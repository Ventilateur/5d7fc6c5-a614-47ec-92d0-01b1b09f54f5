package database

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

type Config struct {
	ConnectionString string
	DBName           string
}

type DB struct {
	users    *mongo.Collection
	userAuth *mongo.Collection
}

func (db *DB) Init(c context.Context, dbConf Config) error {
	client, err := mongo.Connect(c, options.Client().ApplyURI(dbConf.ConnectionString))
	if err != nil {
		return err
	}
	database := client.Database(dbConf.DBName)
	db.users = database.Collection("users")
	db.userAuth = database.Collection("user_auth")
	return err
}

func (db *DB) GetUserHash(c context.Context, userId string) (string, error) {
	userAuth := UserAuthDAO{}
	result := db.userAuth.FindOne(c, bson.M{"id": userId})
	if err := result.Err(); err != nil {
		return "", err
	}
	return userAuth.Hash, result.Decode(&userAuth)
}

func (db *DB) SaveUserHash(c context.Context, userAuth UserAuthDAO) error {
	upsert := true
	_, err := db.userAuth.ReplaceOne(
		c,
		bson.M{"id": userAuth.Id},
		userAuth,
		&options.ReplaceOptions{Upsert: &upsert},
	)
	return err
}

func (db *DB) SaveUserInfo(c context.Context, userInfo UserInfoDAO) error {
	upsert := true
	_, err := db.users.ReplaceOne(
		c,
		bson.M{"id": userInfo.Id},
		userInfo,
		&options.ReplaceOptions{Upsert: &upsert},
	)
	return err
}

func (db *DB) GetUser(c context.Context, userId string) (UserInfoDAO, error) {
	userInfo := UserInfoDAO{}
	result := db.users.FindOne(c, bson.M{"id": userId})
	if err := result.Err(); err != nil {
		return userInfo, err
	}
	return userInfo, result.Decode(&userInfo)
}
