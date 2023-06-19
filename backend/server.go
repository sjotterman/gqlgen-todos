package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/joho/godotenv"
	"github.com/sjotterman/gqlgen-todos/graph"
	"github.com/sjotterman/gqlgen-todos/sqlc/pg"

	_ "github.com/lib/pq"
)

const defaultPort = "8080"

type envVars struct {
	port           string
	dbUsername     string
	dbPassword     string
	dbHost         string
	dbName         string
	clerkSecretKey string
}

func getEnvVars() (envVars, error) {
	envVars := envVars{}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	envVars.port = port
	username, exists := os.LookupEnv("DB_USERNAME")
	if !exists {
		return envVars, fmt.Errorf("DB_USERNAME not set")
	}
	envVars.dbUsername = username
	password, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		return envVars, fmt.Errorf("DB_PASSWORD not set")
	}
	envVars.dbPassword = password
	host, exists := os.LookupEnv("DB_HOST")
	if !exists {
		return envVars, fmt.Errorf("DB_HOST not set")
	}
	envVars.dbHost = host
	dbname, exists := os.LookupEnv("DB_NAME")
	if !exists {
		return envVars, fmt.Errorf("DB_NAME not set")
	}
	envVars.dbName = dbname
	clerkSecretKey, exists := os.LookupEnv("CLERK_SECRET_KEY")
	if !exists {
		return envVars, fmt.Errorf("CLERK_SECRET_KEY not set")
	}
	envVars.clerkSecretKey = clerkSecretKey
	return envVars, nil
}

var allowedOrigins = []string{"http://localhost:3000", "https://menus.otterman.dev"}

func checkCookieHandler(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("in (no-op) checkCookieHandler")
		cookie, err := r.Cookie("__session")
		fmt.Println("cookie", cookie)
		fmt.Println("cookie err", err)
		next.ServeHTTP(w, r)
	}
}

func authHandler(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		fmt.Println("in authHandler")
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		token := ""
		if len(splitToken) >= 2 {
			token = splitToken[1]
		} else {
			fmt.Println("no token")
		}
		fmt.Println("token:", token)
		_, ok := ctx.Value(clerk.ActiveSessionClaims).(*clerk.SessionClaims)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		next.ServeHTTP(w, r)
	}
}

// https://stackoverflow.com/a/64064331
func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		fmt.Println("origin:", origin)
		for _, allowedOrigin := range allowedOrigins {
			if allowedOrigin == origin {
				fmt.Println("allowed origin", allowedOrigin)
				w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
			}
		}
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if r.Method == "OPTIONS" {
			http.Error(w, "No Content", http.StatusNoContent)
			return
		}

		next(w, r)
	}
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	envVars, err := getEnvVars()
	if err != nil {
		log.Fatal(err)
	}
	client, err := clerk.NewClient(envVars.clerkSecretKey)
	if err != nil {
		log.Fatal(err)
	}
	injectActiveSession := clerk.WithSession(client)
	dbString := fmt.Sprintf("postgres://%s:%s@%s/%s", envVars.dbUsername, envVars.dbPassword, envVars.dbHost, envVars.dbName)
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
	withCors := CORS(withCookie)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", withCors)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", envVars.port)
	log.Fatal(http.ListenAndServe(":"+envVars.port, nil))
}
