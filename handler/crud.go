package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Ventilateur/dataimpact-test/database"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func (h *Handler) CreateUser(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}

	dec := json.NewDecoder(src)

	// read open bracket
	if _, err = dec.Token(); err != nil {
		return err
	}

	// while the array contains values
	for dec.More() {
		user := CreateUserRequest{}
		if err = dec.Decode(&user); err != nil {
			return err
		}

		// Should be inside a transaction
		if err = h.saveCredentials(c.Request().Context(), user.Id, user.Password); err != nil {
			return err
		}
		if err = h.db.SaveUserInfo(c.Request().Context(), user.ToUserInfoDAO()); err != nil {
			return err
		}
		if err = h.fs.Save(user.Id, user.Data); err != nil {
			return err
		}
	}

	// read closing bracket
	_, err = dec.Token()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, jsonMsg("OK"))
}

func (h *Handler) saveCredentials(c context.Context, userId, password string) error {
	hash, err := generateHash(password)
	if err != nil {
		return err
	}
	return h.db.SaveUserHash(c, database.UserAuthDAO{
		Id:   userId,
		Hash: hash,
	})
}

func (h *Handler) GetUser(c echo.Context) error {
	userId := c.Param("id")
	userInfoDAO, err := h.db.GetUser(c.Request().Context(), userId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(http.StatusNotFound, jsonMsg(fmt.Sprintf("User %s not found", userId)))
		}
		return err
	}
	response := GetUserResponse{}
	response.FromUserInfoDAO(userInfoDAO)

	data, err := h.fs.Get(userId)
	if err != nil {
		return err
	}
	response.Data = data

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) ListUsers(c echo.Context) error {
	userInfoMap, err := h.db.ListUsers(c.Request().Context())
	if err != nil {
		return err
	}
	userInfoList := make([]GetUserResponse, 0, len(userInfoMap))
	for userId, userInfoDAO := range userInfoMap {
		userInfo := GetUserResponse{}
		userInfo.FromUserInfoDAO(userInfoDAO)

		data, err := h.fs.Get(userId)
		if err != nil {
			return err
		}
		userInfo.Data = data

		userInfoList = append(userInfoList, userInfo)
	}

	return c.JSON(http.StatusOK, userInfoList)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	userId := c.Param("id")
	count, err := h.db.DeleteUser(c.Request().Context(), userId)
	if err != nil {
		return err
	}
	if count == 0 {
		return c.JSON(http.StatusNotFound, jsonMsg(fmt.Sprintf("User %s not found", userId)))
	}

	if err := h.fs.Delete(userId); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, jsonMsg("Deleted"))
}
