package moose

import (
	"log"
	"net/http"
)

type Middleware struct {
	moose bool
}

// Middleware is a struct that has a ServeHTTP method
func NewMiddleware() *Middleware {
	return &Middleware{true}
}

// The middleware handler
func (l *Middleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Log moose status
	log.Printf("MOOSE: %v\n", l.moose)

	// Call the next middleware handler
	next(rw, r)
}
