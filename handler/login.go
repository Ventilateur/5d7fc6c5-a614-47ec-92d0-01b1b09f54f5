package handler

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func (h *Handler) Login(c echo.Context) error {
	login := LoginRequest{}
	if err := c.Bind(&login); err != nil {
		return err
	}

	hash, err := h.db.GetUserHash(c.Request().Context(), login.UserId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.Logger().Info(fmt.Sprintf("User %s does not exist", login.UserId))
			return c.JSON(http.StatusUnauthorized, jsonMsg("Unauthorized"))
		}
		return err
	}

	isAuthorized, err := isAuthorized(login, hash)
	if err != nil {
		return err
	}

	if !isAuthorized {
		c.Logger().Info(fmt.Sprintf("Wrong password for user %s", login.UserId))
		return c.JSON(http.StatusUnauthorized, jsonMsg("Unauthorized"))
	}

	token, err := h.tokenizer.generateToken(login.UserId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, LoginResponse{
		UserId: login.UserId,
		Token:  token,
	})
}
