package middleware

import (
	"fmt"
	"github.com/gorilla/context"
	"net/http"
)

func BasicLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			context := context.Get(r, RequestContextKey).(RequestContext)
			fmt.Printf("RequestId=%v, Method=%v URL=%v\n", context.requestId, r.Method, r.URL)
		}()
		next.ServeHTTP(w, r)
	})
}
