package middleware

import (
	"github.com/davidoram/turbo-octo-avenger/context"
	"log"
	"net/http"
)

//
// Logs some basic information about the request:
// - RequestID
// - URL
// - Method
// - Response
//
func BasicLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			// TODO: Add response
			log.Printf("severity=INFO RequestId=%v, Method=%v URL=%v, Response=\n", context.MustGetRequestId(r), r.Method, r.URL)
		}()
		next.ServeHTTP(w, r)
	})
}
