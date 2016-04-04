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
		fmt.Fprintf(w, "Bafflesnark!")
	})

	n := negroni.Classic()

	// Moose status
	n.Use(moose.NewMiddleware())

	// Handler goes last
	n.UseHandler(mux)

	// Serve
	n.Run(":9030")
}
