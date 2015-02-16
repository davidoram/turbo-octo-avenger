package services

import (
	"github.com/davidoram/turbo-octo-avenger/ipc"
	"github.com/jmoiron/sqlx"
	"github.com/nu7hatch/gouuid"
	"log"
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

func (s *PingService) List(requestId *uuid.UUID, db *sqlx.DB, rv *ipc.ReturnValue) error {

	log.Printf("RequestId=%v Ping::List. db=%v", requestId, db)

	var p PingRow
	err := db.QueryRowx("SELECT message FROM ping LIMIT 1").StructScan(&p)
	if err != nil {
		return err
	}
	log.Printf("RequestId=%v Ping::List query ok", requestId)
	r := ipc.PingResult{p.Message}
	rv.Data = append(rv.Data, &r)
	return nil
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
