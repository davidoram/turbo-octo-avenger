package middleware

import (
	//"github.com/gorilla/context"
	"net/http"
	"time"
)

// RequestTimer sets the startTime & endTime on a request
func RequestTimer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		SetStartTime(r, time.Now())
		defer func() {
			t := time.Now()
			SetEndTime(r, t)
			SetDuration(r, t.Sub(MustGetStartTime(r)))
		}()

		next.ServeHTTP(w, r)
	})
}
