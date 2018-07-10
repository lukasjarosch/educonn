package main

import (
	"github.com/dgrijalva/jwt-go"
	pb "github.com/lukasjarosch/educonn/srv/user/proto/user"
	"time"
)

var (
	key = []byte("wOhijsru7hMD5cPWshoy2fo351IzR0uCAIpFHAL6DiJzClQ1H_NCBEAhMgee9gDiD6QNdFAjwmirgyMbbIcM6w==")
)

type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type TokenService struct {
	repo Respository
}

func (srv *TokenService) Decode(token string) (*CustomClaims, error) {
	// Parse token
	tokenType, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if claims, ok := tokenType.Claims.(*CustomClaims); ok && tokenType.Valid {
		return claims, nil
	}
	return nil, err
}

func (srv *TokenService) Encode(user *pb.User) (string, error) {
	var expireToken = time.Now().Add(time.Hour * 72).Unix()

	claims := CustomClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "educonn.user",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(key)
}
