package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

//temporary directory
var TmpDir = filepath.FromSlash("G:/Repositories/GoFileserverStorage")

func main() {
	// create file server handler
	fs := http.FileServer(http.Dir(TmpDir))

	//handle '/' route
	http.HandleFunc("", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello Golang!")
	})

	//return a '.pdf' file for '/pdf' route
	http.HandleFunc("/pdf", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, filepath.Join(TmpDir, "/files/test.pdf"))
	})

	//return a '.html' file for '/index.html' route
	http.HandleFunc("/index.html", func(req http.ResponseWriter, res *http.Request) {
		http.ServeFile(req, res, filepath.Join(TmpDir, "/index.html"))
	})

	// start HTTP server with `fs` as the default handler
	log.Fatal(http.ListenAndServe(":9000", fs))
}
