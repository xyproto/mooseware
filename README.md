Mooseware
=========

[![Build Status](https://travis-ci.org/xyproto/mooseware.svg?branch=master)](https://travis-ci.org/xyproto/mooseware)

Simple example/skeleton code for writing a Negroni middleware handler.

Middleware example
------------------

~~~ go
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
func (l *Middleware) ServeHTTP(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	// Log moose status
	log.Printf("MOOSE: %v\n", l.moose)

	// Call the next middleware handler
	next(w, req)
}
~~~

Usage
-----

~~~ go
package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/xyproto/mooseware"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Bafflesnark!")
	})

	n := negroni.Classic()

	// Moose status
	n.Use(moose.NewMiddleware())

	// Handler goes last
	n.UseHandler(mux)

	// Serve
	n.Run(":3000")
}
~~~
