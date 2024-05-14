package main

import (
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
)

// GenerateReport creates a table with the stats of the game, Correct answers
// are print in Green, wrong answers are print in Red, not answered words,
// are print in grey
func GenerateReport(words map[string]*Word) {
	var counter int

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Translation", "Word", "Attempts"})
	for _, word := range words {
		counter++
		if word.wrongCounter == 0 {
			t.AppendRow(table.Row{
				defaultAnswerStyle.Render(strconv.Itoa(counter)),
				defaultAnswerStyle.Render(word.translation),
				defaultAnswerStyle.Render(word.text),
				defaultAnswerStyle.Render(strconv.Itoa(word.wrongCounter))})
		} else if word.guessed {
			t.AppendRow(table.Row{
				correctAnswerStyle.Render(strconv.Itoa(counter)),
				correctAnswerStyle.Render(word.translation),
				correctAnswerStyle.Render(word.text),
				correctAnswerStyle.Render(strconv.Itoa(word.wrongCounter))})
		} else {
			t.AppendRow(table.Row{
				wrongAnswerStyle.Render(strconv.Itoa(counter)),
				wrongAnswerStyle.Render(word.translation),
				wrongAnswerStyle.Render(word.text),
				wrongAnswerStyle.Render(strconv.Itoa(word.wrongCounter))})
		}
		t.AppendSeparator()
	}
	t.SetStyle(table.StyleRounded)
	t.Render()
}
