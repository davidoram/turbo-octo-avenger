package main

import (
	"fmt"
	"github.com/davidoram/turbo-octo-avenger/middleware"
	"github.com/davidoram/turbo-octo-avenger/services"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
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

	myHandler := http.HandlerFunc(service.List)
	listChain := alice.New(
		context.ClearHandler,
		middleware.RequestId,
		middleware.BasicLog).Then(myHandler)
	path := fmt.Sprintf("/v%d/%s", service.Version(), service.Name())
	router.Handler("GET", path, listChain)

	// showApppHandler := http.HandlerFunc(service.Show)
	// showChain := alice.New(middleware.BasicLog).Then(showAppHandler)
	// router.GET(fmt.Sprintf("/v%d/%s/:id", service.Version(), service.Name()), showChain)
}
