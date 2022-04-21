package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/guiwoo/exercise_backend/model"
	"github.com/guiwoo/exercise_backend/user"
)

var service = model.DbHandler()

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createUser": &graphql.Field{
			Type:        graphql.String, // return type for this field
			Description: "Create a new user",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name, _ := p.Args["name"].(string)
				email, _ := p.Args["email"].(string)
				password, _ := p.Args["password"].(string)
				service.InsertOne(&model.User_Type{Name: name, Email: email, Password: password})
				return "ok", nil
			},
		},
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
