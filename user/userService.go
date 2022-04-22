package user

import (
	"errors"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/guiwoo/exercise_backend/model"
)

type MutationReturn struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

var service = model.DB_Handler()

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
	service.Insert(&model.User_Type{Name: name, Email: email, Password: password})
	return &MutationReturn{Ok: true, Error: "nil"}, nil
}
