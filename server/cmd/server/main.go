package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/dkalytovskyi/architecture-lab-2/server/db"
)

var httpPortNumber = flag.Int("p", 8080, "HTTP port number")

func NewDbConnection() (*sql.DB, error) {
	conn := &db.Connection{
		DbName:     "greenhouse",
		User:       "postgres",
		Password:   "Denis.23291",
		Host:       "localhost",
		DisableSSL: true,
	}
	return conn.Open()
}

func main() {
	flag.Parse()

	if server, err := ComposeApiServer(HttpPortNumber(*httpPortNumber)); err == nil {
		go func() {
			log.Println("Starting chat server...")

			err := server.Start()
			if err == http.ErrServerClosed {
				log.Printf("HTTP server stopped")
			} else {
				log.Fatalf("Cannot start HTTP server: %s", err)
			}
		}()

		sigChannel := make(chan os.Signal, 1)
		signal.Notify(sigChannel, os.Interrupt)
		<-sigChannel

		if err := server.Stop(); err != nil && err != http.ErrServerClosed {
			log.Printf("Error stopping the server: %s", err)
		}
	} else {
		log.Fatalf("Cannot initialize chat server: %s", err)
	}
}
