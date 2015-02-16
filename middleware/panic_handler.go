package middleware

import (
	// "github.com/davidoram/turbo-octo-avenger/ipc"
	"net/http"
)

//
// Create a handler to handle panics
// When a panic occurs, log an error, return 500 with the panic data
//
func PanicHandler(chain http.Handler) http.Handler {
	var f http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e := recover(); e != nil {
				// TODO Add logging
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		chain.ServeHTTP(w, r)
	}
	return f
}
