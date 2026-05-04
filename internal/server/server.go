package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	HTTP   *http.Server
}

func NewServer(logg *log.Logger) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.MainHandle)
	mux.HandleFunc("/upload", handlers.UploadHandle)

	serv := &http.Server{
		Addr:         ":8080",          //Addr — используйте порт 8080.
		Handler:      mux,              //Handler — передайте ваш http-роутер.
		ErrorLog:     logg,             //ErrorLog — передайте ваш логгер.
		ReadTimeout:  5 * time.Second,  //ReadTimeout — таймаут для чтения. 5 секунд.
		WriteTimeout: 10 * time.Second, //WriteTimeout — таймаут для записи. 10 секунд.
		IdleTimeout:  15 * time.Second, //IdleTimeout — таймаут ожидания следующего запроса. 15 секунд.

	}

	return &Server{
		Logger: logg,
		HTTP:   serv,
	}
}
