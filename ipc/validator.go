package ipc

import (
	"github.com/nu7hatch/gouuid"
	"log"
	"net/http"
)

//
// Validate the URL params according to what is expected from the service
// Append to e if errors occur
//
func ValidateURLParams(requestId *uuid.UUID, r *http.Request, s Service, m ServiceMethod, rv *ReturnValue) {
	log.Printf("level=DEBUG RequestID=%v ValidateURLParams", requestId)
}
