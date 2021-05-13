package handler

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Tokenizer struct {
	TokenDuration time.Duration
	SigningKey    []byte
}

func generateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func isAuthorized(login LoginRequest, hash string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(login.Password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (a *Tokenizer) generateToken(userId string) (string, error) {
	now := time.Now()
	claims := jwt.StandardClaims{
		ExpiresAt: now.Add(a.TokenDuration).Unix(),
		IssuedAt:  now.Unix(),
		Id:        userId,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.SigningKey)
}
