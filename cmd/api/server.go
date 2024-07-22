package main

import (
	"database/sql"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"graphblog/graph"
	"graphblog/internal/article"
	"graphblog/internal/config"
	"graphblog/internal/middleware"
	"graphblog/internal/user"
	"log"
	"net/http"
)

func main() {

	logConf := zap.NewProductionConfig()
	logConf.OutputPaths = []string{
		`env/full/app.log`,
	}

	logger, err := logConf.Build()
	if err != nil {
		log.Fatal(err)
	}

	logger.Info(`boostrapping application`)

	config.Bootstrap()

	dbConnString := fmt.Sprintf(
		"postgresql://blog_user:devpass@%s:%s/blogsite?sslmode=disable",
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
	)
	db, err := sql.Open(`postgres`, dbConnString)
	if err != nil {
		panic(err)
	}

	// dependency injection
	userDAO := user.NewPostgresDAO(db, logger)
	userSvc := user.NewDefaultService(userDAO, logger)

	articleDAO := article.NewPostgresDAO(db, logger)
	articleSvc := article.NewDefaultService(articleDAO, logger)

	resolver := graph.NewResolver(articleSvc, userSvc, logger)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	mux := http.NewServeMux()

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	middle := middleware.NewMiddleware(logger)

	serveHost := fmt.Sprintf("0.0.0.0:%s", config.AppConfig.Port)
	server := http.Server{
		Addr:    serveHost,
		Handler: middle.Wrap(mux),
	}

	log.Printf("connect to %s/ for GraphQL playground", config.AppConfig.Host)
	err = server.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
	}
}
