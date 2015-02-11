package middleware

import (
	"github.com/davidoram/turbo-octo-avenger/context"
	"net/http"
	"time"
)

//
// RequestTimer sets the StartTime, EndTime, and Duration in the context
// The StartTime is set before making calls to other middleare, and the EndTime
// and Duration are set when all the Middleware returns
//
func RequestTimer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		context.SetStartTime(r, time.Now())
		defer func() {
			t := time.Now()
			context.SetEndTime(r, t)
			context.SetDuration(r, t.Sub(context.MustGetStartTime(r)))
		}()

		next.ServeHTTP(w, r)
	})
}
