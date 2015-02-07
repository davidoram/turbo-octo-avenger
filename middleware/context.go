package middleware

//
// Define all the keys for values we store in the request context
//
type key int

// Holds a UUID that identifies the request
const RequestIdKey key = 0
