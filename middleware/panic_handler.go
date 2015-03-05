package middleware

import (
	// "github.com/davidoram/turbo-octo-avenger/ipc"
	"net/http"
)

//
// Handle panics, logs the error, return 500
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
