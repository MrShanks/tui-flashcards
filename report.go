package main

import (
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
)

func GenerateReport(words map[string]*Word) {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Translation", "Word", "Attempts"})
	var counter int
	for _, word := range words {
		counter++
		if word.wrongCounter > 1 {
			t.AppendRow(table.Row{
				wrongAnswerStyle.Render(strconv.Itoa(counter)),
				wrongAnswerStyle.Render(word.translation),
				wrongAnswerStyle.Render(word.text),
				wrongAnswerStyle.Render(strconv.Itoa(word.wrongCounter))})
		} else {
			t.AppendRow(table.Row{
				correctAnswerStyle.Render(strconv.Itoa(counter)),
				correctAnswerStyle.Render(word.translation),
				correctAnswerStyle.Render(word.text),
				correctAnswerStyle.Render(strconv.Itoa(word.wrongCounter))})
		}
		t.AppendSeparator()
	}
	t.SetStyle(table.StyleRounded)
	t.Render()
}
