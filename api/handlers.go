package api

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/MrShanks/tui-flashcards/game"
)

var (
	Words      []*game.Word
	Counter    int
	Score      = 0
	Iterations = 10
)

func StartGame(w http.ResponseWriter, r *http.Request) {
	homepage, err := template.ParseFiles("templates/index.html")
	tmpl := template.Must(homepage, err)
	if len(Words) == 0 {
		Restart(w, r)
	}
	tmpl.Execute(w, Words[Counter].Text)
}

func Restart(w http.ResponseWriter, r *http.Request) {
	Words = game.PickRandomWordsSlice(Iterations)
	Counter = 0
	homepage, err := template.ParseFiles("templates/index.html")
	tmpl := template.Must(homepage, err)
	tmpl.Execute(w, Words[Counter].Text)
	Score = 0
}

func Next(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("%s method is not allowed", r.Method)
		return
	}
	if len(Words) == 0 {
		scr, err := template.ParseFiles("templates/card_score.html")
		tmplScore := template.Must(scr, err)
		tmplScore.Execute(w, fmt.Sprintf("Congrats your score is: %d", Score))
		return
	}
	normal, err := template.ParseFiles("templates/card_normal.html")
	tmpl := template.Must(normal, err)
	Counter++
	if Counter >= len(Words) {
		Counter = 0
	}
	tmpl.Execute(w, Words[Counter].Text)
}

func Prev(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("%s method is not allowed", r.Method)
		return
	}
	if len(Words) == 0 {
		scr, err := template.ParseFiles("templates/card_score.html")
		tmplScore := template.Must(scr, err)
		tmplScore.Execute(w, fmt.Sprintf("Congrats your score is: %d", Score))
		return
	}
	normal, err := template.ParseFiles("templates/card_normal.html")
	tmpl := template.Must(normal, err)
	Counter--
	if Counter < 0 {
		Counter = len(Words) - 1
	}
	tmpl.Execute(w, Words[Counter].Text)
}

func Guess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("%s method is not allowed", r.Method)
		return
	}
	if len(Words) == 0 {
		scr, err := template.ParseFiles("templates/card_score.html")
		tmplScore := template.Must(scr, err)
		tmplScore.Execute(w, fmt.Sprintf("Congrats your score is: %d", Score))
		return
	}

	wrong, err := template.ParseFiles("templates/card_wrong.html")
	tmplWrong := template.Must(wrong, err)

	normal, err := template.ParseFiles("templates/card_normal.html")
	tmplNormal := template.Must(normal, err)

	if r.FormValue("word") == Words[Counter].Translation {
		Score++
		if Counter == len(Words)-1 {
			Words = Words[:len(Words)-1]
			Counter = 0
			if len(Words) == 0 {
				scr, err := template.ParseFiles("templates/card_score.html")
				tmplScore := template.Must(scr, err)
				tmplScore.Execute(w, fmt.Sprintf("Congrats your score is: %d", Score))
				return
			}
			tmplNormal.Execute(w, Words[Counter].Text)
		} else {
			Words = append(Words[:Counter], Words[Counter+1:]...)
			tmplNormal.Execute(w, Words[Counter].Text)
		}
	} else {
		Score--
		tmplWrong.Execute(w, Words[Counter])
	}

}
