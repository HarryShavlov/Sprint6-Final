package main

import (
	"log"
	"os"

	"sprint6/internal/server"
)

func main() {
	logg := log.New(os.Stdout, "info: ", log.LstdFlags)

	serv := server.NewServer(logg)

	err := serv.HTTP.ListenAndServe()
	if err != nil {
		logg.Fatal(err)
	}
}
