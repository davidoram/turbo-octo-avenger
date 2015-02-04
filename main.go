package main

import (
	"fmt"
	// "github.com/davidoram/turbo-octo-avenger/services"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// Any struct that is marshalled to/from JSON
type JSONStruct struct{}

// Defines service type
type ApiServiceImpl interface {
	Name() string
	Version() int
	//registerRoutes(router httprouter.Router)

	List(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Show(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	//Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	//Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	//Delete(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
}

// func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	fmt.Fprint(w, "Welcome!\n")
// }
//
// func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
// }

func main() {
	router := httprouter.New()
	//router.GET("/", Index)
	//router.GET("/hello/:name", Hello)
	pingSvc := new(PingService)
	register(pingSvc, router)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func register(service ApiServiceImpl, router *httprouter.Router) {
	router.GET(fmt.Sprintf("/v%d/%s", service.Version(), service.Name()), service.List)
	router.GET(fmt.Sprintf("/v%d/%s/:id", service.Version(), service.Name()), service.Show)
}

// -----------------------------------------------------------------------------
//
// Ping Service
//

type PingService struct {
}

func (s *PingService) Name() string {
	return "ping"
}

func (s *PingService) Version() int {
	return 1
}

func (s *PingService) List(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "{ 'data': [1,2,3] }")
}

func (s *PingService) Show(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "{ 'id': %v  }", p.ByName("id"))
}
