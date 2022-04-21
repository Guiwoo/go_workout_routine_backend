package graphql

import (
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	schema "github.com/guiwoo/exercise_backend/graphql/schema"
)

var PORT string

func GraphQLStart(port string) {
	PORT = port

	graphqlHandler := handler.New(&handler.Config{
		Schema:     &schema.UserSchema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	http.Handle("/graphql", graphqlHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
