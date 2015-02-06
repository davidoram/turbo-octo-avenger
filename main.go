package main

import (
	"fmt"
	"github.com/davidoram/turbo-octo-avenger/services"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// Any struct that is marshalled to/from JSON
type JSONStruct struct{}

func main() {
	router := httprouter.New()

	services := []services.Service{new(services.PingService)}

	for _, service := range services {
		register(service, router)
	}

	log.Fatal(http.ListenAndServe(":8080", router))
}

func register(service services.Service, router *httprouter.Router) {
	router.GET(fmt.Sprintf("/v%d/%s", service.Version(), service.Name()), service.List)
	router.GET(fmt.Sprintf("/v%d/%s/:id", service.Version(), service.Name()), service.Show)
}
