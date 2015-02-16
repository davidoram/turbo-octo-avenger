package middleware

import (
	"github.com/davidoram/turbo-octo-avenger/context"
	"github.com/davidoram/turbo-octo-avenger/ipc"
	"github.com/davidoram/turbo-octo-avenger/util"
	"log"
	"net/http"
)

//
// Is responsible for translating HTTP middleware semantics to IPC semantics as follows:
// - Performs validation on the URL params, Headers, Body values
// - Parses the URL params, Headers, Body, and context into structures
// - Constructs an return value structure containing code, errors, body
// - Calls the service method
// - Validate the return value
// - Encodes the return values into Headers, Body, return code
//
func ServiceCaller(s ipc.Service, m ipc.ServiceMethod) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestId := context.MustGetRequestId(r)
		var rv = ipc.ReturnValue{RequestID: requestId}

		// If a panic occurs then encode it
		defer func() {
			if e := recover(); e != nil {
				err := e.(error)
				rv.Errors = append(rv.Errors, &err)
				rv.StatusCode = http.StatusInternalServerError
				log.Printf("level=ERROR RequestId=%v Source=%v Action='Caught Panic' Error=%v", requestId, s, e)
			}
			// Turn the ReturnValue into HTTP JSON response
			w.WriteHeader(rv.StatusCode)
			if len(rv.Errors) == 0 {
				w.Write(util.MustMarshalJSON(rv.Data))
			} else {
				w.Write(util.MustMarshalJSON(rv.Errors))
			}
		}()

		ipc.ValidateURLParams(requestId, r, s, m, &rv)
		if len(rv.Errors) == 0 {
			err := s.List(requestId, context.MustGetDB(r), &rv)
			if err != nil {
				rv.Errors = append(rv.Errors, &err)
				rv.StatusCode = http.StatusInternalServerError
				log.Printf("level=ERROR RequestId=%v Source=%v Action='Service returned error' Error=%v", requestId, s, err)
			}
		}

	})
}
