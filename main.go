package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/MrShanks/tui-flashcards/game"
	_ "github.com/lib/pq"
)

var words []*game.Word
var counter int

func homepage(w http.ResponseWriter, r *http.Request) {
	homepage, err := template.ParseFiles("index.html")
	tmpl := template.Must(homepage, err)
	tmpl.Execute(w, fmt.Sprintf("%d: %s", counter, words[counter].Text))
}

func guess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("%s method is not allowed", r.Method)
		return
	}
	if r.FormValue("word") == words[counter].Translation {
		counter++
	}
	if counter >= 10 {
		counter = 0
	}
	w.Write([]byte(fmt.Sprintf("%d: %s", counter, words[counter].Text)))
}

func main() {
	deck, err := game.LoadWords()
	if err != nil {
		log.Printf("Couldn't load words: %s", err)
	}
	words = deck[:10]

	http.HandleFunc("/", homepage)
	http.HandleFunc("/guess", guess)

	log.Fatal(http.ListenAndServe(":8080", nil))
	// game.Repl()
}
