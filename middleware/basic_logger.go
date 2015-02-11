package middleware

import (
	"log"
	"net/http"
	"time"
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
			log.Printf("RequestId=%v, Method=%v URL=%v, MillisecDuration=%d\n", MustGetRequestId(r), r.Method, r.URL, MustGetDuration(r)*time.Millisecond)
		}()
		next.ServeHTTP(w, r)
	})
}
