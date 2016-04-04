Mooseware
=========

[![Build Status](https://travis-ci.org/xyproto/mooseware.svg?branch=master)](https://travis-ci.org/xyproto/mooseware)

Simple example/skeleton code for writing a Negroni middleware handler.

Usage
-----

~~~ go
package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/xyproto/middleskel"
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
	n.Run(":9030")
}
~~~
