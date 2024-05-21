package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/MrShanks/tui-flashcards/game"
	_ "github.com/lib/pq"
)

var (
	words      []*game.Word
	counter    int
	score      = 0
	iterations = 10
)

func StartGame(w http.ResponseWriter, r *http.Request) {
	homepage, err := template.ParseFiles("templates/index.html")
	tmpl := template.Must(homepage, err)
	if len(words) == 0 {
		Restart(w, r)
	}
	tmpl.Execute(w, words[counter].Text)
}

func Restart(w http.ResponseWriter, r *http.Request) {
	words = game.PickRandomWordsSlice(iterations)
	counter = 0
	homepage, err := template.ParseFiles("templates/index.html")
	tmpl := template.Must(homepage, err)
	tmpl.Execute(w, words[counter].Text)
	score = 0
}

func Next(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("%s method is not allowed", r.Method)
		return
	}
	if len(words) == 0 {
		scr, err := template.ParseFiles("templates/card_score.html")
		tmplScore := template.Must(scr, err)
		tmplScore.Execute(w, fmt.Sprintf("Congrats your score is: %d", score))
		return
	}
	normal, err := template.ParseFiles("templates/card_normal.html")
	tmpl := template.Must(normal, err)
	counter++
	if counter >= len(words) {
		counter = 0
	}
	tmpl.Execute(w, words[counter].Text)
}

func Prev(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("%s method is not allowed", r.Method)
		return
	}
	if len(words) == 0 {
		scr, err := template.ParseFiles("templates/card_score.html")
		tmplScore := template.Must(scr, err)
		tmplScore.Execute(w, fmt.Sprintf("Congrats your score is: %d", score))
		return
	}
	normal, err := template.ParseFiles("templates/card_normal.html")
	tmpl := template.Must(normal, err)
	counter--
	if counter < 0 {
		counter = len(words) - 1
	}
	tmpl.Execute(w, words[counter].Text)
}

func Guess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("%s method is not allowed", r.Method)
		return
	}
	if len(words) == 0 {
		scr, err := template.ParseFiles("templates/card_score.html")
		tmplScore := template.Must(scr, err)
		tmplScore.Execute(w, fmt.Sprintf("Congrats your score is: %d", score))
		return
	}

	wrong, err := template.ParseFiles("templates/card_wrong.html")
	tmplWrong := template.Must(wrong, err)

	normal, err := template.ParseFiles("templates/card_normal.html")
	tmplNormal := template.Must(normal, err)

	if r.FormValue("word") == words[counter].Translation {
		score++
		if counter == len(words)-1 {
			words = words[:len(words)-1]
			counter = 0
			if len(words) == 0 {
				scr, err := template.ParseFiles("templates/card_score.html")
				tmplScore := template.Must(scr, err)
				tmplScore.Execute(w, fmt.Sprintf("Congrats your score is: %d", score))
				return
			}
			tmplNormal.Execute(w, words[counter].Text)
		} else {
			words = append(words[:counter], words[counter+1:]...)
			tmplNormal.Execute(w, words[counter].Text)
		}
	} else {
		score--
		tmplWrong.Execute(w, words[counter].Text)
	}

}
func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", StartGame)
	mux.HandleFunc("/guess", Guess)
	mux.HandleFunc("/next", Next)
	mux.HandleFunc("/prev", Prev)
	mux.HandleFunc("/restart", Restart)
}

func NewServer() http.Handler {
	mux := http.NewServeMux()
	AddRoutes(mux)
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

	words = game.PickRandomWordsSlice(iterations)

	log.Fatal(httpServer.ListenAndServeTLS("certs/localhost+2.pem", "certs/localhost+2-key.pem"))
	// game.Repl()
}
