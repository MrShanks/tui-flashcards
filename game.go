package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var score, counter int
var wrongWords = make(map[string]*Word)
var exit = false
var iterations = 10

func GuessTheWord(wordMap map[string]*Word, scanner *bufio.Scanner) {
	for text, word := range wordMap {
		clearScreen()
		fmt.Printf("Write the translation %d/%d\t\t\t\tScore: %d\n", counter, len(wordMap), score)
		fmt.Println(normal.Render(text))
		fmt.Print("> ")

		scanner.Scan()
		input := scanner.Text()
		args := strings.Fields(input)
		if input == "stop" {
			exit = true
			return
		}

		if word.translation != input || len(args) == 0 {
			clearScreen()

			word.wrongCounter++
			wrongWords[text] = word

			fmt.Println(wrong.Render(fmt.Sprintf("%s : %s", text, input)))
			fmt.Println("Expected: ", correctAnswerStyle.Render(word.translation))
			printExample(word.example)
			fmt.Println("Press enter to continue")

			scanner.Scan()

			clearScreen()
			continue
		}

		if word.translation == input {
			clearScreen()

			word.wrongCounter++
			word.guessed = true
			score++

			fmt.Println(correct.Render(fmt.Sprintf("%s : %s", text, input)))
			fmt.Println("Great that was the right answer")
			printExample(word.example)
			fmt.Println("Press enter to continue")

			scanner.Scan()

			clearScreen()
		}
	}
}
func GuessTheWrongWords(wordMap map[string]*Word, scanner *bufio.Scanner) {
	for text, word := range wordMap {
		clearScreen()
		fmt.Printf("Write the translation \t\t\t\tScore: %d\n", score)
		fmt.Println(normal.Render(text))
		fmt.Print("> ")

		scanner.Scan()
		input := scanner.Text()
		args := strings.Fields(input)
		if input == "stop" {
			exit = true
			return
		}

		if word.translation != input || len(args) == 0 {
			clearScreen()

			word.wrongCounter++
			wrongWords[text] = word

			fmt.Println(wrong.Render(fmt.Sprintf("%s : %s", text, input)))
			fmt.Println("Expected: ", correctAnswerStyle.Render(word.translation))
			printExample(word.example)
			fmt.Println("Press enter to continue")

			scanner.Scan()

			clearScreen()
			continue
		}

		if word.translation == input {
			clearScreen()

			word.wrongCounter++
			word.guessed = true

			fmt.Println(correct.Render(fmt.Sprintf("%s : %s", text, input)))
			fmt.Println("Great that was the right answer")
			printExample(word.example)
			fmt.Println("Press enter to continue")

			scanner.Scan()

			delete(wordMap, text)

			clearScreen()
		}
	}
}

func NewGame(iterations int) {
	scanner := bufio.NewScanner(os.Stdin)
	wordMap := pickRandomWords(iterations)

	bestScore, err := ReadBestScore("bestScore.txt")
	if err != nil {
		fmt.Println(err)
	}

	clearScreen()
	fmt.Println("Hi Welcome back to your language training")
	fmt.Println("Ready?")
	scanner.Scan()

	GuessTheWord(wordMap, scanner)

	// start the second loop to answer the words you didn't get right in the first run
	for {
		if len(wrongWords) == 0 || exit {
			break
		}
		GuessTheWrongWords(wrongWords, scanner)
	}

	// Update the score in case you have scored better than your best
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
	score = 0
	counter = 0
}
