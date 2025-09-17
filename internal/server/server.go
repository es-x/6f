package server

import (
	"log"
	"net/http"
	"time"

	"github.com/es-x/6f/internal/handlers"
)

type Server struct {
	Logger     log.Logger
	HttpServer *http.Server
}

func NewServer(l *log.Logger) *Server {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.MainHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	myServer := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger:     *l,
		HttpServer: myServer,
	}
}

// s := &http.Server{
// 	Addr:           ":8080",
// 	Handler:        myHandler,
// 	ReadTimeout:    10 * time.Second,
// 	WriteTimeout:   10 * time.Second,
// 	MaxHeaderBytes: 1 << 20,
// }
// log.Fatal(s.ListenAndServe())
