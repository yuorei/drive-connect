package main

import (
	"drive-connect-bff/directive"
	"drive-connect-bff/graph/generated"
	"log"
	"net/http"
	"os"

	"drive-connect-bff/graph/resolver"
	"drive-connect-bff/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	resolver := resolver.NewResolver()
	c := generated.Config{Resolvers: resolver}
	c.Directives.Auth = directive.Auth
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", middleware.AuthMiddleware(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
