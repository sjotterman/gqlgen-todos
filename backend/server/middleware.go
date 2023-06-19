package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/clerkinc/clerk-sdk-go/clerk"
)

var allowedOrigins = []string{"http://localhost:3000", "https://menus.otterman.dev"}

func CheckCookieHandler(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("in (no-op) checkCookieHandler")
		cookie, err := r.Cookie("__session")
		fmt.Println("cookie", cookie)
		fmt.Println("cookie err", err)
		next.ServeHTTP(w, r)
	}
}

func AuthHandler(next http.Handler) http.HandlerFunc {
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
