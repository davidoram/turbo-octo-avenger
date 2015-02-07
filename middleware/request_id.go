package middleware

import (
	"github.com/gorilla/context"
	"github.com/nu7hatch/gouuid"
	"net/http"
)

// Sets the 'RequestId' UUID in the context
func RequestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		u4, err := uuid.NewV4()
		if err != nil {
			panic(err)
		}

		context.Set(r, RequestIdKey, u4)

		next.ServeHTTP(w, r)
	})
}
