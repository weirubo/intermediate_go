package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// jwt-go

type Claims struct {
	ID       int64
	Username string
	jwt.StandardClaims
}

var secret = "golang"

func main() {
	token, err := GenerateToken()
	if err != nil {
		fmt.Printf("GenerateToken() err: %v\n", err)
		return
	}
	fmt.Printf("token: %v\n", token)

	claims, err := ParseToken(token)
	if err != nil {
		fmt.Printf("ParseToken() err: %v\n", err)
		return
	}
	fmt.Printf("claims: %+v\n", claims)
}

func GenerateToken() (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(300 * time.Second)
	issuer := "frank"
	claims := Claims{
		ID:       10001,
		Username: "frank",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
