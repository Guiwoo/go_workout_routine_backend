package user

import (
	"github.com/graphql-go/graphql"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var MutationReturnType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Return",
	Fields: graphql.Fields{
		"ok": &graphql.Field{
			Type: graphql.Boolean,
		},
		"error": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var LoginReturnType = graphql.NewObject(graphql.ObjectConfig{
	Name: "LoginReturnType",
	Fields: graphql.Fields{
		"ok": &graphql.Field{
			Type: graphql.Boolean,
		},
		"error": &graphql.Field{
			Type: graphql.String,
		},
		"token": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var FindUserReturnType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FindUsersReturnType",
	Fields: graphql.Fields{
		"ok": &graphql.Field{
			Type: graphql.Boolean,
		},
		"error": &graphql.Field{
			Type: graphql.String,
		},
		"users": &graphql.Field{
			Type: graphql.NewList(UserType),
		},
	},
})

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
	},
})
