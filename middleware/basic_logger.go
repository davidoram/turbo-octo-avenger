package middleware

import (
	"github.com/davidoram/turbo-octo-avenger/context"
	"log"
	"net/http"
	_ "time"
)

//
// Logs some basic information about the request:
// - RequestID
// - URL
// - Method
// - Milliseconds duration
//
func BasicLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			log.Printf("RequestId=%v, Method=%v URL=%v, Response=\n", context.MustGetRequestId(r), r.Method, r.URL)
		}()
		next.ServeHTTP(w, r)
	})
}
