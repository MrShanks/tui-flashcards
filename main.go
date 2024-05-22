package main

import (
	"log"
	"net/http"
	"time"

	"github.com/MrShanks/tui-flashcards/api"
	_ "github.com/lib/pq"
)

func main() {
	api.Game = api.NewGame()
	mux := api.NewMux()

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	log.Fatal(httpServer.ListenAndServeTLS("certs/localhost+2.pem", "certs/localhost+2-key.pem"))
	// game.Repl()
}
