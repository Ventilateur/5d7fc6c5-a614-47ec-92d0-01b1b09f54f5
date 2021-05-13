package database

type UserAuthDAO struct {
	Id   string `bson:"id"`
	Hash string `bson:"hash"`
}

type FriendDAO struct {
	Id   int    `bson:"id"`
	Name string `bson:"name"`
}

type UserInfoDAO struct {
	Id         string      `bson:"id"`
	IsActive   bool        `bson:"isActive"`
	Balance    string      `bson:"balance"`
	Age        int         `bson:"age"`
	Name       string      `bson:"name"`
	Gender     string      `bson:"gender"`
	Company    string      `bson:"company"`
	Email      string      `bson:"email"`
	Phone      string      `bson:"phone"`
	Address    string      `bson:"address"`
	About      string      `bson:"about"`
	Registered string      `bson:"registered"`
	Latitude   float64     `bson:"latitude"`
	Longitude  float64     `bson:"longitude"`
	Tags       []string    `bson:"tags"`
	Friends    []FriendDAO `bson:"friends"`
}
