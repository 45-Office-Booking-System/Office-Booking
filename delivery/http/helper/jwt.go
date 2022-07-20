package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type MyClaims struct {
	UserID int    `json:"userID"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	jwt.StandardClaims
}

type GoJWT struct {
}

func NewGoJWT() *GoJWT {
	return &GoJWT{}
}

func (j *GoJWT) CreateTokenJWT(userID int, email string, name string, phone string) (string, error) {
	claims := MyClaims{
		UserID: userID,
		Email:  email,
		Name:   name,
		Phone:  phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("220220"))
}
