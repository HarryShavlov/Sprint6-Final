package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logg := log.New(os.Stdout, "info: ", log.LstdFlags)

	serv := server.NewServer(logg)

	err := serv.HTTP.ListenAndServe()
	if err != nil {
		logg.Fatal(err)
	}
}
