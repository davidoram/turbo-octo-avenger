package middleware

import (
	"github.com/gorilla/context"
	"github.com/nu7hatch/gouuid"
	"net/http"
	"time"
)

//
// Define all the keys for values we store in the request context
//
type key int

// Holds a UUID that identifies the request
const RequestContextKey key = 0

type RequestContext struct {
	requestId *uuid.UUID // UUID uniuqe to this request

	startTime time.Time     // When the request started
	endTime   time.Time     // When the request finished
	duration  time.Duration // Request duration

}

// Creates the initial values in the context
func ContextSetup(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		u4, err := uuid.NewV4()
		if err != nil {
			panic(err)
		}

		rc := RequestContext{
			requestId: u4,
		}
		context.Set(r, RequestContextKey, rc)

		next.ServeHTTP(w, r)
	})
}
