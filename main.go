package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const ver = "0.1.1"

func main() {
	dir, _ := os.Getwd()
	port := flag.String("port", "8000", "specify port number to listen to")
	test := flag.Bool("test", false, "Print working directory and exit")
	flag.Parse()
	log.Printf("Trying to bind to port %s\n", *port)
	log.Printf("qhttp %s serving %s\n", ver, dir)

	if *test == true {
		os.Exit(0)
	}
	handler := logRequest(noCache(http.FileServer(http.Dir(dir))))
	panic(http.ListenAndServe(fmt.Sprintf(":%s", *port), handler))
}
