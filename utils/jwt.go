package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//PUT TO .ENV
var secretKey []byte = []byte("test")

type UserInfo struct {
	ID uint `json:"id"`
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

func VerifyToken(token string) (uint, error) {
	verifiedToken, err := jwt.ParseWithClaims(token, &CustomClaimsExample{}, func(jwtToken *jwt.Token) (interface{}, error) {
		return []byte("test"), nil
	})

	if claims, ok := verifiedToken.Claims.(*CustomClaimsExample); ok && verifiedToken.Valid {
		return claims.ID, nil
	} else {
		return 0, err
	}
}
