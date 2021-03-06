package user

import (
	"github.com/graphql-go/graphql"
)

var CreateUser = &graphql.Field{
	Type:        MutationReturnType, // return type for this field
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
	Resolve: CreateUserService,
}

var LoginUser = &graphql.Field{
	Type:        LoginReturnType, // return type for this field
	Description: "Create a new user",
	Args: graphql.FieldConfigArgument{
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: LoginUserService,
}

var EditUser = &graphql.Field{
	Type:        MutationReturnType,
	Description: "Edit a User",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: EditUserService,
}

// Query
var SearchUser = &graphql.Field{
	Type:        SearchUserReturnType,
	Description: "Search users",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: SearchUserService,
}

var FindUserById = &graphql.Field{
	Type:        FindOneUserReturnType,
	Description: "Find Users by id",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: FindUserByIdService,
}
