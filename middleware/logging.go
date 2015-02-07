package middleware

import (
	"fmt"
	"net/http"
)

func BasicLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("BEFORE %v : %v\n", r.Method, r.URL)
		next.ServeHTTP(w, r)
		fmt.Printf("AFTER %v : %v\n", r.Method, r.URL)
	})
}
