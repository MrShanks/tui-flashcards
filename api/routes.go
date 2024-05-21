package api

import "net/http"

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", StartGame)
	mux.HandleFunc("/guess", Guess)
	mux.HandleFunc("/next", Next)
	mux.HandleFunc("/prev", Prev)
	mux.HandleFunc("/restart", Restart)
}
