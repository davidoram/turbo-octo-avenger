package main

import (
	"fmt"
	"github.com/davidoram/turbo-octo-avenger/middleware"
	"github.com/davidoram/turbo-octo-avenger/services"
	"github.com/gorilla/context"
	"github.com/jmoiron/sqlx"
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

	configureLogToSyslog()

	router := httprouter.New()

	services := []services.Service{new(services.PingService)}

	for _, service := range services {
		register(service, router)
	}

	log.Fatal(http.ListenAndServe(":8080", router))
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

func register(service services.Service, router *httprouter.Router) {

	db1 := sqlx.MustConnect("postgres", "postgres://davidoram:@localhost/turbo-octo-avenger-development?sslmode=disable")
	e := db1.Ping()
	if e != nil {
		panic(e)
	}
	log.Printf("Database connection setup ok : %v", db1)

	myHandler := http.HandlerFunc(service.List)
	listChain := alice.New(
		context.ClearHandler,
		middleware.PanicHandler,
		middleware.RequestIDInjector,
		middleware.BasicLogger,
		middleware.RequestTimer,
		middleware.DatabaseConnectionInjector(db1)).Then(myHandler)
	path := fmt.Sprintf("/v%d/%s", service.Version(), service.Name())
	router.Handler("GET", path, listChain)

	// showApppHandler := http.HandlerFunc(service.Show)
	// showChain := alice.New(middleware.BasicLog).Then(showAppHandler)
	// router.GET(fmt.Sprintf("/v%d/%s/:id", service.Version(), service.Name()), showChain)
}
