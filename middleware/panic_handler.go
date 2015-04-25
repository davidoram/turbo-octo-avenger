package middleware

import (
	"encoding/json"
	"github.com/davidamitchell/turbo-octo-avenger/context"
	"github.com/davidamitchell/turbo-octo-avenger/services"
	"log"
	"net/http"
	"runtime"
)

//
// Handle panics, logs the error, return 500
//
func PanicHandler(chain http.Handler) http.Handler {
	var f http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e := recover(); e != nil {
				// Take care because of error inside this error handler
				buf := make([]byte, 1<<16)
				runtime.Stack(buf, true)
				log.Printf("severity=ERROR RequestID=%v e=%v stack=%s", context.GetRequestId(r), e, buf)
				w.WriteHeader(http.StatusInternalServerError)
				var response = services.NewPingResponse(context.GetRequestId(r))
				response.Errors = append(response.Errors, services.Error{services.ErrorCodeInternalServerError, "Internal error - see logs"})
				b, err := json.Marshal(response)
				if err == nil {
					w.Write(b)
				} else {
					log.Printf("severity=ERROR RequestID=%v e=%v message='error encoding the error!'", context.GetRequestId(r), err)
				}

			}
		}()
		chain.ServeHTTP(w, r)
	}
	return f
}
