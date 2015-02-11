package context

import (
	//"database/sql"
	"github.com/gorilla/context"
	"github.com/jmoiron/sqlx"
	"github.com/nu7hatch/gouuid"
	//"log"
	"net/http"
	"time"
)

//
// Define all the keys for values we store in the request context
//
const (
	RequestIDKey = "toa:request_id"
	StartTimeKey = "toa:start_time"
	EndTimeKey   = "toa:end_time"
	DurationKey  = "toa:duration"
	DBKey        = "toa:db"
)

func MustGetRequestId(r *http.Request) *uuid.UUID {
	rv := GetRequestId(r)
	if rv == nil {
		panic("RequestID not yet set in context")
	}
	return rv
}

func GetRequestId(r *http.Request) *uuid.UUID {
	if rv := context.Get(r, RequestIDKey); rv != nil {
		return rv.(*uuid.UUID)
	}
	return nil
}

func SetRequestID(r *http.Request, val *uuid.UUID) {
	context.Set(r, RequestIDKey, val)
}

func MustGetStartTime(r *http.Request) time.Time {
	if rv := context.Get(r, StartTimeKey); rv != nil {
		return rv.(time.Time)
	}
	panic("StartTime not yet set in context")
}

func SetStartTime(r *http.Request, val time.Time) {
	context.Set(r, StartTimeKey, val)
}

func MustGetEndTime(r *http.Request) time.Time {
	if rv := context.Get(r, EndTimeKey); rv != nil {
		return rv.(time.Time)
	}
	panic("EndTime not yet set in context")
}

func SetEndTime(r *http.Request, val time.Time) {
	context.Set(r, EndTimeKey, val)
}

func MustGetDuration(r *http.Request) time.Duration {
	if rv := context.Get(r, DurationKey); rv != nil {
		return rv.(time.Duration)
	}
	panic("Duration not yet set in context")
}

func SetDuration(r *http.Request, val time.Duration) {
	context.Set(r, DurationKey, val)
}

func MustGetDB(r *http.Request) *sqlx.DB {
	rv := GetDB(r)
	if rv == nil {
		panic("DB not yet set in context")
	}
	return rv
}

func GetDB(r *http.Request) *sqlx.DB {
	if rv := context.Get(r, DBKey); rv != nil {
		return rv.(*sqlx.DB)
	}
	return nil
}

func SetDB(r *http.Request, val *sqlx.DB) {
	context.Set(r, DBKey, val)
}
