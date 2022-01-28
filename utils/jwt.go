package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//PUT TO .ENV
var secretKey []byte = []byte("test")

type UserInfo struct {
	ID uint
}

type CustomClaimsExample struct {
	*jwt.StandardClaims
	TokenType string
	UserInfo
}

func CreateToken(id uint) (string, error) {

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomClaimsExample{
		&jwt.StandardClaims{

			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		},
		"level1",
		UserInfo{id}})

	token, err := t.SignedString(secretKey)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return token, nil
}
