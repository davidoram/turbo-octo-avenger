package middleware

import (
	"github.com/davidoram/turbo-octo-avenger/context"
	"github.com/nu7hatch/gouuid"
	"net/http"
)

//
// Injects a unique RequestId into the context
// Should be the first middleware to be invoked as all other middleware can expect this to
// be there. The RequestID is used to track all log messages across the layers.
//
func RequestIDInjector(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		u4, err := uuid.NewV4()
		if err != nil {
			panic(err)
		}
		context.SetRequestID(r, u4.String())

		next.ServeHTTP(w, r)
	})
}
