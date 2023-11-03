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

	// CORSのハンドラを設定
	corsHandler := allowCORS(srv)

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", middleware.AuthMiddleware(corsHandler))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// CORSを許可するハンドラ
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ヘッダーの設定
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

		// オプションリクエストの場合は即座にレスポンスを返す
		if r.Method == "OPTIONS" {
			return
		}

		// 通常のリクエストの場合はハンドラを呼び出す
		h.ServeHTTP(w, r)
	})
}
