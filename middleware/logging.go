package middleware

import (
	"github.com/gorilla/context"
	"log"
	"net/http"
	"time"
)

func BasicLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			context := context.Get(r, RequestContextKey).(RequestContext)
			log.Printf("RequestId=%v, Method=%v URL=%v, MillisecDuration=%d\n", context.requestId, r.Method, r.URL, context.duration*time.Millisecond)
		}()
		next.ServeHTTP(w, r)
	})
}
