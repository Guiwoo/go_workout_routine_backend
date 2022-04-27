package user

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/graphql-go/graphql"
	"github.com/guiwoo/exercise_backend/model"
	"golang.org/x/crypto/bcrypt"
)

const Cost int = 10

type MutationReturn struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

type LoginReturn struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
	Token string `json:"token"`
}

type Claims struct {
	Email string `json:"email"`
	Id    int64  `json:"id"`
	jwt.StandardClaims
}

var (
	service = model.DB_Handler()
	JwtKey  = []byte("HolyWak")
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), Cost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateToken(email string, id int64) string {
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

var CreateUserService = func(p graphql.ResolveParams) (interface{}, error) {
	name, _ := p.Args["name"].(string)
	email, _ := p.Args["email"].(string)
	password, _ := p.Args["password"].(string)

	var user = model.User_Type{Email: email}
	has, err := service.Get(&user)
	if err != nil {
		log.Fatal(err)
	}
	if has {
		return &MutationReturn{Ok: false, Error: "email has already taken"}, errors.New("email has already taken")
	}
	password, _ = hashPassword(password)
	service.Insert(&model.User_Type{Name: name, Email: email, Password: password})
	return &MutationReturn{Ok: true, Error: "nil"}, nil
}

var LoginUserService = func(p graphql.ResolveParams) (interface{}, error) {
	var findUsers []model.User_Type
	email, _ := p.Args["email"].(string)
	password, _ := p.Args["password"].(string)
	err := service.In("email", email).Find(&findUsers)
	if err != nil {
		log.Fatal(err)
	}
	if len(findUsers) < 1 {
		return LoginReturn{Ok: false, Error: "could not find the email", Token: ""}, errors.New("could not find the email")
	}
	if ok := checkPasswordHash(password, findUsers[0].Password); !ok {
		return LoginReturn{Ok: false, Error: "password is not correct", Token: ""}, errors.New("password is not correct")
	}
	token := generateToken(email, findUsers[0].ID)
	return LoginReturn{Ok: true, Error: "lala", Token: token}, nil
}

var EditUserService = func(p graphql.ResolveParams) (interface{}, error) {
	name, _ := p.Args["name"].(string)
	// password, _ := p.Args["password"].(string)
	holy := p.Info.RootValue.(map[string]interface{})
	jwt, ok := holy["jwt"]
	if !ok {
		return &MutationReturn{Ok: false, Error: "Need to Login Account First"}, nil
	}
	ok, id := verifyingToken(jwt.(string))
	if !ok {
		return &MutationReturn{Ok: false, Error: "Need to Login Account First"}, nil
	}
	if name != "" {
		user := new(model.User_Type)
		user.Name = name
		affected, err := service.ID(id).Update(user)
		if err != nil {
			log.Fatal(err)
		}
		if affected < 1 {
			return &MutationReturn{Ok: false, Error: "Update failed"}, nil
		}
	}
	return &MutationReturn{Ok: true, Error: "nil"}, nil
}
