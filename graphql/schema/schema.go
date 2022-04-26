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
		"user": &graphql.Field{
			Type:        user.UserType,
			Description: "Get Single User",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return user.User{ID: 1, Name: "guiwoo", Email: "pbk@gmail.com", Password: "123"}, nil
			},
		},
	},
})

var UserSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})
