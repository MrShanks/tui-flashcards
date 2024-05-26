package api

import "net/http"

func NewMux() http.Handler {
	mux := http.NewServeMux()
	Game.AddRoutes(mux)
	return mux
}

func (g *GameState) AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", g.Homepage)
	mux.HandleFunc("/settings", g.Settings)
	mux.HandleFunc("/save-settings", SaveSettings)
	mux.HandleFunc("/start", g.StartGame)
	mux.HandleFunc("/guess", g.Guess)
	mux.HandleFunc("/next", g.Next)
	mux.HandleFunc("/prev", g.Prev)
	mux.HandleFunc("/restart", g.Restart)
}
