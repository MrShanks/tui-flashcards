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
var score = 0
var iterations = 10

func homepage(w http.ResponseWriter, r *http.Request) {
	homepage, err := template.ParseFiles("index.html")
	tmpl := template.Must(homepage, err)
	if len(words) == 0 {
		restart(w, r)
	}
	tmpl.Execute(w, words[counter].Text)
}

func restart(w http.ResponseWriter, r *http.Request) {
	words = game.PickRandomWordsSlice(iterations)
	counter = 0
	homepage, err := template.ParseFiles("index.html")
	tmpl := template.Must(homepage, err)
	tmpl.Execute(w, words[counter].Text)
	score = 0
}

func next(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("%s method is not allowed", r.Method)
		return
	}
	if len(words) == 0 {
		scr, err := template.ParseFiles("card_score.html")
		tmplScore := template.Must(scr, err)
		tmplScore.Execute(w, fmt.Sprintf("Congrats your score is: %d", score))
		return
	}
	normal, err := template.ParseFiles("card_normal.html")
	tmpl := template.Must(normal, err)
	counter++
	if counter >= len(words) {
		counter = 0
	}
	tmpl.Execute(w, words[counter].Text)
}

func prev(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("%s method is not allowed", r.Method)
		return
	}
	if len(words) == 0 {
		scr, err := template.ParseFiles("card_score.html")
		tmplScore := template.Must(scr, err)
		tmplScore.Execute(w, fmt.Sprintf("Congrats your score is: %d", score))
		return
	}
	normal, err := template.ParseFiles("card_normal.html")
	tmpl := template.Must(normal, err)
	counter--
	if counter < 0 {
		counter = len(words) - 1
	}
	tmpl.Execute(w, words[counter].Text)
}

func guess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("%s method is not allowed", r.Method)
		return
	}
	if len(words) == 0 {
		scr, err := template.ParseFiles("card_score.html")
		tmplScore := template.Must(scr, err)
		tmplScore.Execute(w, fmt.Sprintf("Congrats your score is: %d", score))
		return
	}

	wrong, err := template.ParseFiles("card_wrong.html")
	tmplWrong := template.Must(wrong, err)

	normal, err := template.ParseFiles("card_normal.html")
	tmplNormal := template.Must(normal, err)

	if r.FormValue("word") == words[counter].Translation {
		score++
		if counter == len(words)-1 {
			words = words[:len(words)-1]
			counter = 0
			if len(words) == 0 {
				scr, err := template.ParseFiles("card_score.html")
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

func main() {
	words = game.PickRandomWordsSlice(iterations)

	http.HandleFunc("/", homepage)
	http.HandleFunc("/guess", guess)
	http.HandleFunc("/next", next)
	http.HandleFunc("/prev", prev)
	http.HandleFunc("/restart", restart)

	log.Fatal(http.ListenAndServe(":8080", nil))
	// game.Repl()
}
