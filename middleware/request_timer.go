package middleware

import (
	"github.com/davidoram/turbo-octo-avenger/context"
	"log"
	"net/http"
	"time"
)

//
// RequestTimer logs the duration of each call
//
func RequestTimer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			millis := int64((time.Now().Sub(start)) / time.Millisecond)
			log.Printf("RequestId=%v, Duration_ms=%v\n", context.MustGetRequestId(r), millis)
		}()

		next.ServeHTTP(w, r)
	})
}
