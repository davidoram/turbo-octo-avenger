package main

import (
	_ "fmt"
	"github.com/davidamitchell/turbo-octo-avenger/middleware"
	"github.com/davidamitchell/turbo-octo-avenger/services"
	"github.com/davidamitchell/turbo-octo-avenger/services/userservice"
	"github.com/gorilla/context"
	_ "github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	_ "github.com/lib/pq"
	"log"
	"log/syslog"
	"net/http"
)

const (
	AppLogTag = "turbo-octo-avenger"
)

func main() {

	//configureLogToSyslog()

	router := httprouter.New()

	registerPingService(router)
	registerUserService(router)
	port := ":8080"
	log.Printf("severity=INFO port=%v", port)
	log.Fatal(http.ListenAndServe(port, router))
}

// Redirect the global logger to use syslog
func configureLogToSyslog() {
	// Configure logger to write to the syslog
	logwriter, e := syslog.New(syslog.LOG_NOTICE, AppLogTag)
	if e != nil {
		panic(e)
	}
	log.SetOutput(logwriter)
}

func registerPingService(router *httprouter.Router) {

	verb := "GET"
	path := "/v1/ping"
	listChain := alice.New(
		context.ClearHandler,
		middleware.PanicHandler,
		middleware.RequestIDInjector,
		middleware.BasicLogger,
		middleware.RequestTimer).Then(services.PingServiceListHandler())
	router.Handler(verb, path, listChain)
	log.Printf("severity=INFO action='added route' verb=%s path='%s'", verb, path)
}

func registerUserService(router *httprouter.Router) {

	verb := "POST"
	path := "/users"
	listChain := alice.New(
		context.ClearHandler,
		middleware.PanicHandler,
		middleware.RequestIDInjector,
		middleware.BasicLogger,
		middleware.RequestTimer).Then(userservice.UserServiceInsertHandler())
	router.Handler(verb, path, listChain)
	log.Printf("severity=INFO action='added route' verb=%s path='%s'", verb, path)
}
