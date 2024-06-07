package main

import (
	"database/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/james-cathcart/golog"
	_ "github.com/lib/pq"
	"graphblog/graph"
	"graphblog/internal/article"
	"graphblog/internal/user"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {

	logger := golog.NewLogger(golog.NewNativeLogger(`[ main ] `))
	logger.Info(`application loading`)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := sql.Open(`postgres`, `postgresql://blog_user:devpass@localhost:5432/blogsite?sslmode=disable`)
	if err != nil {
		panic(err)
	}

	// dependency injection
	userDAO := user.NewPostgresDAO(db)
	userSvc := user.NewDefaultService(userDAO)

	articleDAO := article.NewPostgresDAO(db)
	articleSvc := article.NewDefaultService(articleDAO)

	resolver := graph.NewResolver(articleSvc, userSvc)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
