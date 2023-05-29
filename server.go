package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/sjotterman/gqlgen-todos/graph"
	"github.com/sjotterman/gqlgen-todos/sqlc/pg"

	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	username, exists := os.LookupEnv("DB_USERNAME")
	if !exists {
		log.Fatal("DB_USERNAME not set")
	}
	password, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		log.Fatal("DB_PASSWORD not set")
	}
	host, exists := os.LookupEnv("DB_HOST")
	if !exists {
		log.Fatal("DB_HOST not set")
	}
	dbname, exists := os.LookupEnv("DB_NAME")
	if !exists {
		log.Fatal("DB_NAME not set")
	}

	dbString := fmt.Sprintf("postgres://%s:%s@%s/%s", username, password, host, dbname)
	db, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queries := pg.New(db)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Queries: queries}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
