package server

import (
	"flag"
	"log"
	"net/http"
)

var (
	listen = flag.String("listen", ":8083", "listen address")
	dir    = "./web/"
)

func StartStaticServer() {
	flag.Parse()
	log.Printf("listening on %q...", *listen)
	err := http.ListenAndServe(*listen, http.FileServer(http.Dir(dir)))
	log.Fatalln(err)
}
