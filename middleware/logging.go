package middleware

import (
	//"github.com/gorilla/context"
	"log"
	"net/http"
	"time"
)

func BasicLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			log.Printf("RequestId=%v, Method=%v URL=%v, MillisecDuration=%d\n", MustGetRequestId(r), r.Method, r.URL, MustGetDuration(r)*time.Millisecond)
		}()
		next.ServeHTTP(w, r)
	})
}
