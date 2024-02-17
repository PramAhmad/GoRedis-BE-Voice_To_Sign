package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("7EKhqyxC3j-zXmBs0VOTqq-6Kk2lYA3G2bFqoLS3fLTa8zioEyxAP6Xbjv4vyWVVN5pDdRd9QiPkFWk5Lj5WQA")

type JWTClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJwt(username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString(jwtKey)

	return

}

func VerifyJwt(tokenString string) (claims *JWTClaim, err error) {
	claims = &JWTClaim{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return
	}
	if !token.Valid {
		return claims, err
	}
	return claims, nil

}
