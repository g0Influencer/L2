package main

import (
	"L2/develop/dev11/server"
	"flag"
	"log"
	"net/http"
)

func main() {
	P := flag.String("p", "80", "port")
	flag.Parse()

	mux := http.NewServeMux()

	handler := server.NewHandler()

	handler.Register(mux)

	muxWithLogger := server.Logging(mux)

	log.Printf("Server is listening on: %s\n", *P)
	err := http.ListenAndServe("localhost:"+*P, muxWithLogger)
	if err != nil {
		log.Fatal(err)
	}

}
