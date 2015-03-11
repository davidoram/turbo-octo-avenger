package context

import (
	"github.com/gorilla/context"
	"net/http"
)

//
// Define all the keys for values we store in the request context
//
const (
	RequestIDKey = "toa:request_id"
)

func MustGetRequestId(r *http.Request) string {
	if rv := context.Get(r, RequestIDKey); rv != nil {
		return rv.(string)
	}
	panic("RequestID not yet set in context")
}

func GetRequestId(r *http.Request) string {
	if rv := context.Get(r, RequestIDKey); rv != nil {
		return rv.(string)
	}
	return "RequestID not yet set in context"
}

func SetRequestID(r *http.Request, val string) {
	context.Set(r, RequestIDKey, val)
}
