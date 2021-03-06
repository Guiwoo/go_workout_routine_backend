package graphql

import (
	"context"
	"log"
	"net/http"
	"strings"

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
		RootObjectFn: func(ctx context.Context, r *http.Request) map[string]interface{} {
			holy := r.Header["Jwt"]
			TheMap := make(map[string]interface{})
			TheMap["jwt"] = strings.Join(holy, "")
			return TheMap
		},
	})

	http.Handle("/graphql", graphqlHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
