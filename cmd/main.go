package main

import (
	"log"
	"os"

	"github.com/es-x/6f/internal/server"
)

func main() {
	logger := log.New(os.Stdout, `serv`, log.LstdFlags|log.Lshortfile)

	s := server.NewServer(logger)

	err := s.HttpServer.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
