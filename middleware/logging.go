package middleware

import (
	"fmt"
	"github.com/gorilla/context"
	"net/http"
)

func BasicLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		fmt.Printf("RequestId=%v, Method=%v URL=%v\n", context.Get(r, RequestIdKey), r.Method, r.URL)
	})
}
