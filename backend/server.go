package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/joho/godotenv"
	"github.com/sjotterman/gqlgen-todos/graph"
	"github.com/sjotterman/gqlgen-todos/server"
	"github.com/sjotterman/gqlgen-todos/sqlc/pg"

	_ "github.com/lib/pq"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	envVars, err := server.GetEnvVars()
	if err != nil {
		log.Fatal(err)
	}
	client, err := clerk.NewClient(envVars.ClerkSecretKey)
	if err != nil {
		log.Fatal(err)
	}
	injectActiveSession := clerk.WithSession(client)
	dbString := fmt.Sprintf("postgres://%s:%s@%s/%s", envVars.DbUsername, envVars.DbPassword, envVars.DbHost, envVars.DbName)
	if(strings.Contains(envVars.DbHost, "0.0.0.0")) { 
		fmt.Println("Local DB, disabling sslmode")
		dbString += "?sslmode=disable"
	} 
	db, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queries := pg.New(db)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Queries: queries}}))
	withAuth := server.AuthHandler(srv)
	serverWithSession := injectActiveSession(withAuth)
	withCookie := server.CheckCookieHandler(serverWithSession)
	withCors := server.CORS(withCookie)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", withCors)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", envVars.Port)
	log.Fatal(http.ListenAndServe(":"+envVars.Port, nil))
}
