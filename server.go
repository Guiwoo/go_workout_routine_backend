package main

import (
	"os"

	gql "github.com/guiwoo/exercise_backend/graphql"
	"github.com/guiwoo/exercise_backend/model"
)

const defaultPort = "8080"

func main() {
	model.DBOpen()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	gql.GraphQLStart(port)
}
