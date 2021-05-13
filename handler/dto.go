package handler

import "github.com/Ventilateur/dataimpact-test/database"

type LoginRequest struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

type Friend struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CreateUserRequest struct {
	Id         string   `json:"id"`
	Password   string   `json:"password"`
	IsActive   bool     `json:"isActive"`
	Balance    string   `json:"balance"`
	Age        int      `json:"age"`
	Name       string   `json:"name"`
	Gender     string   `json:"gender"`
	Company    string   `json:"company"`
	Email      string   `json:"email"`
	Phone      string   `json:"phone"`
	Address    string   `json:"address"`
	About      string   `json:"about"`
	Registered string   `json:"registered"`
	Latitude   float64  `json:"latitude"`
	Longitude  float64  `json:"longitude"`
	Tags       []string `json:"tags"`
	Friends    []Friend `json:"friends"`
	Data       string   `json:"data"`
}

func (cur *CreateUserRequest) ToUserInfoDAO() database.UserInfoDAO {
	userInfoDAO := database.UserInfoDAO{
		Id:         cur.Id,
		IsActive:   cur.IsActive,
		Balance:    cur.Balance,
		Age:        cur.Age,
		Name:       cur.Name,
		Gender:     cur.Gender,
		Company:    cur.Company,
		Email:      cur.Email,
		Phone:      cur.Phone,
		Address:    cur.Address,
		About:      cur.About,
		Registered: cur.Registered,
		Latitude:   cur.Latitude,
		Longitude:  cur.Longitude,
		Tags:       make([]string, len(cur.Tags)),
		Friends:    make([]database.FriendDAO, len(cur.Friends)),
	}
	for i, f := range cur.Friends {
		userInfoDAO.Friends[i] = database.FriendDAO{
			Id:   f.Id,
			Name: f.Name,
		}
	}
	copy(userInfoDAO.Tags, cur.Tags)

	return userInfoDAO
}

type GetUserResponse struct {
	Id         string   `json:"id"`
	IsActive   bool     `json:"isActive"`
	Balance    string   `json:"balance"`
	Age        int      `json:"age"`
	Name       string   `json:"name"`
	Gender     string   `json:"gender"`
	Company    string   `json:"company"`
	Email      string   `json:"email"`
	Phone      string   `json:"phone"`
	Address    string   `json:"address"`
	About      string   `json:"about"`
	Registered string   `json:"registered"`
	Latitude   float64  `json:"latitude"`
	Longitude  float64  `json:"longitude"`
	Tags       []string `json:"tags"`
	Friends    []Friend `json:"friends"`
	Data       string   `json:"data"`
}

func (gur *GetUserResponse) FromUserInfoDAO(userInfoDAO database.UserInfoDAO) {
	gur.Id = userInfoDAO.Id
	gur.IsActive = userInfoDAO.IsActive
	gur.Balance = userInfoDAO.Balance
	gur.Age = userInfoDAO.Age
	gur.Name = userInfoDAO.Name
	gur.Gender = userInfoDAO.Gender
	gur.Company = userInfoDAO.Company
	gur.Email = userInfoDAO.Email
	gur.Phone = userInfoDAO.Phone
	gur.Address = userInfoDAO.Address
	gur.About = userInfoDAO.About
	gur.Registered = userInfoDAO.Registered
	gur.Latitude = userInfoDAO.Latitude
	gur.Longitude = userInfoDAO.Longitude
	gur.Tags = make([]string, len(userInfoDAO.Tags))
	gur.Friends = make([]Friend, len(userInfoDAO.Friends))
	copy(gur.Tags, userInfoDAO.Tags)
	for i, f := range userInfoDAO.Friends {
		gur.Friends[i] = Friend{
			Id:   f.Id,
			Name: f.Name,
		}
	}
}
