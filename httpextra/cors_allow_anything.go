package httpextra

import (
	"net/http"
	"strings"
)

// CorsAllowAnythingMiddleware allows any request from anywhere
func CorsAllowAnythingMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.Header().Add("Access-Control-Allow-Headers", "*")
			w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

			if r.Method == http.MethodOptions {
				return
			}

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}

type CustomCors struct {
	AllowOrigin  []string // use ["*"] for all
	AllowHeaders []string // use ["*"] for all
	AllowMethods []string
}

func AllMethods() []string {
	return []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
}

// CorsAllowAnythingMiddleware allows any request from anywhere
func CorsCustomMiddleware(config CustomCors) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Access-Control-Allow-Origin", strings.Join(config.AllowOrigin, ", "))
			w.Header().Add("Access-Control-Allow-Headers", strings.Join(config.AllowHeaders, ", "))
			w.Header().Add("Access-Control-Allow-Methods", strings.Join(config.AllowMethods, ", "))

			if r.Method == http.MethodOptions {
				return
			}

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
