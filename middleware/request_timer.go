package middleware

import (
	"github.com/gorilla/context"
	"net/http"
	"time"
)

// RequestTimer sets the startTime & endTime on a request
func RequestTimer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		context := context.Get(r, RequestContextKey).(RequestContext)

		context.startTime = time.Now()
		defer func() {
			context.endTime = time.Now()
			context.duration = context.endTime.Sub(context.startTime)
		}()

		next.ServeHTTP(w, r)
	})
}
