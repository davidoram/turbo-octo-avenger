package middleware

import (
	"fmt"
	"github.com/davidoram/turbo-octo-avenger/context"
	"log"
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
				log.Printf("level=ERROR RequestId=%v, Panic=%v", context.MustGetRequestId(r), e)
				w.WriteHeader(500)
				fmt.Fprintf(w, "{ 'error': '%v'  }", e)
			}
		}()
		chain.ServeHTTP(w, r)
	}
	return f
}
