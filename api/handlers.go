package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/MrShanks/tui-flashcards/game"
)

var Game *GameState = nil
var CardsNumber int = 10

type GameState struct {
	Words       []*game.Word
	CurrentWord *game.Word
	Counter     int
	Score       int
	WordsLeft   int
}

func NewGame() *GameState {
	words := game.PickRandomWordsSlice(CardsNumber)
	counter := 0
	currentWord := words[counter]
	wordLeft := len(words)
	return &GameState{
		Words:       words,
		CurrentWord: currentWord,
		Counter:     counter,
		Score:       0,
		WordsLeft:   wordLeft,
	}
}

func (g *GameState) Homepage(w http.ResponseWriter, r *http.Request) {
	homepage, err := template.ParseFiles("templates/homepage.html")
	tmpl := template.Must(homepage, err)
	tmpl.Execute(w, r)
}

func (g *GameState) StartGame(w http.ResponseWriter, r *http.Request) {
	g.Words = game.PickRandomWordsSlice(CardsNumber)
	g.Counter = 0
	g.Score = 0
	g.WordsLeft = len(g.Words)
	g.CurrentWord = g.Words[g.Counter]
	game, err := template.ParseFiles("templates/index.html")
	tmpl := template.Must(game, err)
	if len(g.Words) == 0 {
		g.Restart(w, r)
	}
	tmpl.Execute(w, g)
}

func (g *GameState) Restart(w http.ResponseWriter, r *http.Request) {
	*g = *NewGame()
	homepage, err := template.ParseFiles("templates/index.html")
	tmpl := template.Must(homepage, err)
	tmpl.Execute(w, g)
}

func (g *GameState) Next(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("%s method is not allowed", r.Method)
		return
	}
	if len(g.Words) == 0 {
		scr, err := template.ParseFiles("templates/card_score.html")
		tmplScore := template.Must(scr, err)
		tmplScore.Execute(w, g)
		return
	}
	normal, err := template.ParseFiles("templates/card_normal.html")
	tmpl := template.Must(normal, err)
	g.Counter++
	if g.Counter >= len(g.Words) {
		g.Counter = 0
	}
	g.CurrentWord = g.Words[g.Counter]
	tmpl.Execute(w, g)
}

func (g *GameState) Prev(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("%s method is not allowed", r.Method)
		return
	}
	if len(g.Words) == 0 {
		scr, err := template.ParseFiles("templates/card_score.html")
		tmplScore := template.Must(scr, err)
		tmplScore.Execute(w, g)
		return
	}
	normal, err := template.ParseFiles("templates/card_normal.html")
	tmpl := template.Must(normal, err)
	g.Counter--
	if g.Counter < 0 {
		g.Counter = len(g.Words) - 1
	}
	g.CurrentWord = g.Words[g.Counter]
	tmpl.Execute(w, g)
}

func (g *GameState) Guess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("%s method is not allowed", r.Method)
		return
	}
	if len(g.Words) == 0 {
		scr, err := template.ParseFiles("templates/card_score.html")
		tmplScore := template.Must(scr, err)
		tmplScore.Execute(w, g)
		return
	}

	wrong, err := template.ParseFiles("templates/card_wrong.html")
	tmplWrong := template.Must(wrong, err)

	normal, err := template.ParseFiles("templates/card_normal.html")
	tmplNormal := template.Must(normal, err)

	if r.FormValue("word") == g.CurrentWord.Translation {
		g.Score++
		g.WordsLeft--
		if g.Counter == len(g.Words)-1 {
			g.Words = g.Words[:len(g.Words)-1]
			g.CurrentWord = nil
			g.Counter = 0
			if len(g.Words) == 0 {
				game.RegisterStats(int32(g.Score))
				scr, err := template.ParseFiles("templates/card_score.html")
				tmplScore := template.Must(scr, err)
				tmplScore.Execute(w, g)
				return
			}
			tmplNormal.Execute(w, g)
		} else {
			g.Words = append(g.Words[:g.Counter], g.Words[g.Counter+1:]...)
			g.CurrentWord = g.Words[g.Counter]
			tmplNormal.Execute(w, g)
		}
	} else {
		g.Score--
		tmplWrong.Execute(w, g)
	}
}

func (g *GameState) Settings(w http.ResponseWriter, r *http.Request) {
	settings, err := template.ParseFiles("templates/settings.html")
	tmplSettings := template.Must(settings, err)
	tmplSettings.Execute(w, nil)
}

func (g *GameState) Scores(w http.ResponseWriter, r *http.Request) {
	dbscores := game.GetStats()
	scores, err := template.ParseFiles("templates/scores.html")
	tmplSettings := template.Must(scores, err)
	tmplSettings.Execute(w, dbscores)
}

func SaveSettings(w http.ResponseWriter, r *http.Request) {
	settings, err := template.ParseFiles("templates/settings.html")
	tmplSettings := template.Must(settings, err)
	err = r.ParseForm()
	if err != nil {
		fmt.Printf("Unable to parse form: %s", err)
		return
	}

	n := r.FormValue("cards")

	if CardsNumber, err = strconv.Atoi(n); err != nil {
		fmt.Printf("Couldn't convert value: %s to string: %s", n, err)
		return
	}
	tmplSettings.Execute(w, nil)
}
