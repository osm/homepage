package main

import (
	"embed"
	"flag"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

//go:embed index.html
var index []byte

//go:embed static
var staticFS embed.FS

func main() {
	port := flag.String("p", "4001", "listen port")
	flag.Parse()

	files, _ := fs.ReadDir(staticFS, "static")
	for _, f := range files {
		n := f.Name()
		d, _ := staticFS.ReadFile(filepath.Join("static", n))

		http.HandleFunc("/"+n, func(w http.ResponseWriter, req *http.Request) {
			ip := req.Header.Get("X-Real-IP")
			if ip == "" {
				ip = req.RemoteAddr[0:strings.LastIndex(req.RemoteAddr, ":")]
			}
			log.Printf("%s downloaded from %s", n, ip)
			w.Write(d)

		})
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write(index)
	})

	http.ListenAndServe(":"+*port, nil)
}
