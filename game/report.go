package game

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
		if word.WrongCounter == 0 {
			t.AppendRow(table.Row{
				defaultAnswerStyle.Render(strconv.Itoa(counter)),
				defaultAnswerStyle.Render(word.Translation),
				defaultAnswerStyle.Render(word.Text),
				defaultAnswerStyle.Render(strconv.Itoa(word.WrongCounter))})
		} else if word.Guessed {
			if word.WrongCounter > 1 {
				t.AppendRow(table.Row{
					warningAnswerStyle.Render(strconv.Itoa(counter)),
					warningAnswerStyle.Render(word.Translation),
					warningAnswerStyle.Render(word.Text),
					warningAnswerStyle.Render(strconv.Itoa(word.WrongCounter))})
			} else {
				t.AppendRow(table.Row{
					correctAnswerStyle.Render(strconv.Itoa(counter)),
					correctAnswerStyle.Render(word.Translation),
					correctAnswerStyle.Render(word.Text),
					correctAnswerStyle.Render(strconv.Itoa(word.WrongCounter))})
			}
		} else {
			t.AppendRow(table.Row{
				wrongAnswerStyle.Render(strconv.Itoa(counter)),
				wrongAnswerStyle.Render(word.Translation),
				wrongAnswerStyle.Render(word.Text),
				wrongAnswerStyle.Render(strconv.Itoa(word.WrongCounter))})
		}
		t.AppendSeparator()
	}
	t.SetStyle(table.StyleRounded)
	t.Render()
}
