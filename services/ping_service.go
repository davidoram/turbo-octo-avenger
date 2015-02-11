package services

import (
	"fmt"
	"github.com/davidoram/turbo-octo-avenger/middleware"
	//	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

// -----------------------------------------------------------------------------
//
// Ping Service
//

type PingService struct {
}

// Represents database table 'Ping'
type PingRow struct {
	Message string
}

func (s *PingService) Name() string {
	return "ping"
}

func (s *PingService) Version() int {
	return 1
}

func (s *PingService) List(w http.ResponseWriter, r *http.Request) {

	log.Printf("RequestId=%v Ping::List. db=%v", middleware.MustGetRequestId(r), middleware.MustGetDB(r))

	var p PingRow
	db := middleware.MustGetDB(r)
	err := db.QueryRowx("SELECT message FROM ping LIMIT 1").StructScan(&p)
	if err != nil {
		log.Printf("RequestId=%v Ping::List err %v", middleware.MustGetRequestId(r), err)
		panic(err)
	}
	log.Printf("RequestId=%v Ping::List query ok", middleware.MustGetRequestId(r))

	fmt.Fprintf(w, "{ 'pong': '%v'  }", p.Message)
}

// type PingService struct {
// }
//
// func (s *PingService) Name() string {
// 	return "ping"
// }
//
// func (s *PingService) Version() int {
// 	return 1
// }

// func (s *PingService) List(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "{ 'data': [1,2,3] }")
// }
//
// func (s *PingService) Show(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "{ 'id': 1  }")
// }
