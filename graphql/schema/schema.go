package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/guiwoo/exercise_backend/user"
)

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createUser": user.CreateUser,
		"loginUser":  user.LoginUser,
		"editUser":   user.EditUser,
	},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"findUser": user.SearchUser,
		"oneUser":  user.FindUserById,
	},
})

var UserSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})
