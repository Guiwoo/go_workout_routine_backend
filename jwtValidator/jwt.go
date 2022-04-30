package jwtValidator

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/graphql-go/graphql"
)

type Claims struct {
	Email string `json:"email"`
	Id    int64  `json:"id"`
	jwt.StandardClaims
}

type MutationReturn struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

var JwtKey = []byte("HolyWak")

func GenerateToken(email string, id int64) string {
	expirationTime := time.Now().Add(10 * time.Hour)
	claims := &Claims{
		Email: email,
		Id:    id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return ""
	}
	return tokenString
}

func verifyingToken(token string) (bool, int64) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(JwtKey), nil
	})
	fmt.Println(tkn.Claims.(*Claims).Id)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, 0
		}
		return false, 0
	}
	if !tkn.Valid {
		return false, 0
	}
	id := tkn.Claims.(*Claims).Id
	return true, id
}

func JwtValidChecker(p graphql.ResolveParams) (interface{}, error) {
	header := p.Info.RootValue.(map[string]interface{})
	jwt, ok := header["jwt"]
	if !ok {
		return &MutationReturn{Ok: false, Error: "Need to Login Account First"}, nil
	}
	ok, id := verifyingToken(jwt.(string))
	if !ok {
		return &MutationReturn{Ok: false, Error: "Need to Login Account First"}, nil
	}
	return id, nil
}
