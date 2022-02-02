package main

import (
	_ "embed"
	"flag"
	"net/http"
)

//go:embed index.html
var index []byte

//go:embed favicon.png
var favicon []byte

func main() {
	port := flag.String("p", "4001", "listen port")
	flag.Parse()

	http.HandleFunc("/favicon.png", func(w http.ResponseWriter, req *http.Request) {
		w.Write(favicon)

	})
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write(index)
	})

	http.ListenAndServe(":"+*port, nil)
}
