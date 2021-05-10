package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// create file server handler
	fs := http.FileServer(http.Dir("G:/Repositories/GoFileserverStorage"))

	//handle '/' route
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")
		fmt.Fprint(res, "<h1>Go Lang</h1>")
	})

	// handle `/static` route
	http.Handle("/static", fs)

	// start HTTP server with `fs` as the default handler
	log.Fatal(http.ListenAndServe(":9000", fs))
}
