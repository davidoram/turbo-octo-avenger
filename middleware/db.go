package middleware

import (
	//"database/sql"
	//"github.com/gorilla/context"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

func DatabaseConnection(db *sqlx.DB) func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Setting db into context: %v", db)
			SetDB(r, db)
			next.ServeHTTP(w, r)
		})
	}
}
