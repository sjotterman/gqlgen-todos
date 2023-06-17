package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/sjotterman/gqlgen-todos/graph"
	"github.com/sjotterman/gqlgen-todos/sqlc/pg"

	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func checkCookieHandler(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("__session")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		defer func() {
			fmt.Println("leave checkCookieHandler")
		}()
		next.ServeHTTP(w, r)
	}
}

func authHandler(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		fmt.Println("in authHandler")
		_, ok := ctx.Value(clerk.ActiveSessionClaims).(*clerk.SessionClaims)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		next.ServeHTTP(w, r)
	}
}

// func init() {
// 	// loads values from .env into the system
// 	if err := godotenv.Load(); err != nil {
// 		log.Print("No .env file found")
// 	}
// }

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
	clerkSecretKey, exists := os.LookupEnv("CLERK_SECRET_KEY")
	if !exists {
		log.Fatal("CLERK_SECRET_KEY not set")
	}
	client, err := clerk.NewClient(clerkSecretKey)
	if err != nil {
		log.Fatal(err)
	}
	injectActiveSession := clerk.WithSession(client)
	dbString := fmt.Sprintf("postgres://%s:%s@%s/%s", username, password, host, dbname)
	db, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queries := pg.New(db)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Queries: queries}}))
	withAuth := authHandler(srv)
	serverWithSession := injectActiveSession(withAuth)
	withCookie := checkCookieHandler(serverWithSession)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", withCookie)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
