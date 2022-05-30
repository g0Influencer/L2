package main

import (
	"L2/develop/dev11/server"
	"log"
	"net/http"
	"os"
)

func main(){

	mux:=http.NewServeMux()

	handler:=server.NewHandler()

	handler.Register(mux)

	muxWithLogger := server.Logging(mux)

	addr := ":8080"
	func() {
		temp := os.Getenv("PORT")
		if temp != "" {
			addr = temp
		}
	}()


	log.Printf("Server is listening on: %s\n", addr)
	err:= http.ListenAndServe(addr, muxWithLogger)
	if err != nil {
		log.Fatal(err)
	}

}
