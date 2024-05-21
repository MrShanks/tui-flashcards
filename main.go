package main

import (
	"log"
	"net/http"
	"time"

	"github.com/MrShanks/tui-flashcards/api"
	"github.com/MrShanks/tui-flashcards/game"
	_ "github.com/lib/pq"
)

func NewServer() http.Handler {
	mux := http.NewServeMux()
	api.AddRoutes(mux)
	return mux
}

func main() {
	srv := NewServer()

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      srv,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	api.Words = game.PickRandomWordsSlice(api.Iterations)

	log.Fatal(httpServer.ListenAndServeTLS("certs/localhost+2.pem", "certs/localhost+2-key.pem"))
	// game.Repl()
}
