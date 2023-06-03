package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = "80"

type Config struct {
}

func main() {
	app := Config{}

	log.Printf("starting broke service on PORT %s\n", PORT)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: app.routes(),
	}

	// start server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
