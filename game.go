package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

var score, counter int
var wrongWords = make(map[string]*Word)

func startUp(scanner *bufio.Scanner) {
	clearScreen()
	fmt.Println("Hi Welcome back to your language training")
	fmt.Println("Ready?")
	scanner.Scan()
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printExample(example string) {
	fmt.Println("Here is an example:")
	fmt.Println(exampleStyle.Render(example))
}

func translateWords(wordMap map[string]*Word, scanner *bufio.Scanner, repeat bool) {

	for text, word := range wordMap {
		translation := word.translation
		example := word.example
		if !repeat {
			counter++
		}
		clearScreen()
		if repeat {
			fmt.Printf("Write the translation \t\t\t\tScore: %d\n", score)
		} else {
			fmt.Printf("Write the translation %d/%d\t\t\t\tScore: %d\n", counter, len(wordMap), score)
		}
		fmt.Println(normal.Render(text))
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		args := strings.Fields(input)

		if translation != input || len(args) == 0 {
			clearScreen()
			word.wrongCounter++
			wrongWords[text] = word
			fmt.Println(wrong.Render(fmt.Sprintf("%s : %s", text, input)))
			fmt.Println("Expected: ", correctAnswerStyle.Render(translation))
			printExample(example)
			fmt.Println("Press enter to continue")
			scanner.Scan()
			clearScreen()
			continue
		}

		if translation == input {
			clearScreen()
			if !repeat {
				score++
			}
			fmt.Println(correct.Render(fmt.Sprintf("%s : %s", text, input)))
			fmt.Println("Great that was the right answer")
			printExample(example)
			fmt.Println("Press enter to continue")
			scanner.Scan()
			clearScreen()
			if repeat {
				delete(wordMap, text)
			}
		}
	}
}

func NewGame(iterations int) {
	scanner := bufio.NewScanner(os.Stdin)

	words, err := LoadWords()
	if err != nil {
		fmt.Println(err)
	}
	wordMap := pickRandomWords(words, iterations)

	bestScore, err := ReadBestScore("bestScore.txt")
	if err != nil {
		fmt.Println(err)
	}

	startUp(scanner)
	translateWords(wordMap, scanner, false)

	// start the second loop to answer the words you didn't get right
	for {
		if len(wrongWords) == 0 {
			break
		}
		translateWords(wrongWords, scanner, true)
	}

	if score > bestScore {
		UpdateBestScore("bestScore.txt", score, wordMap)
	} else {
		fmt.Printf("Great your final Score is: \n")
		fmt.Println(normal.Render(fmt.Sprintf("%d/%d", score, len(wordMap))))
		fmt.Printf("Don't forget to come back tomorrow!\n")
	}
	fmt.Println("Press enter to see the report")
	scanner.Scan()

	GenerateReport(wordMap)
}

func GenerateReport(words map[string]*Word) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Translation", "Word", "Attempts"})
	var counter int
	for _, word := range words {
		counter++
		if word.wrongCounter >= 1 {
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
