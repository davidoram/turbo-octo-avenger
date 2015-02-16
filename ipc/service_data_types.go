package ipc

// Hold service specific data types in struct form
// The data structures here are understood by the middleware for conversion to & from JSON
// They are used directly by services

// Ping service
type PingResult struct {
	Message string
}
