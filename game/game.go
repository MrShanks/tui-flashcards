package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	score, counter int
	wrongWords     = make(map[string]*Word)
	exit           = false
	iterations     = 10
)

// GuessTheWord starts the first loop of words, score is updated only during this run.
// Not guessed words are stored in the wrongWords map that is later used in the second
// loop
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

		if word.Translation != input || len(args) == 0 {
			clearScreen()

			word.WrongCounter++
			wrongWords[text] = word

			fmt.Println(wrong.Render(fmt.Sprintf("%s : %s", text, input)))
			fmt.Println("Expected: ", correctAnswerStyle.Render(word.Translation))
			printExample(word.Example)
			fmt.Println("Press enter to continue")

			scanner.Scan()

			clearScreen()
			continue
		}

		if word.Translation == input {
			clearScreen()

			word.WrongCounter++
			word.Guessed = true
			score++

			fmt.Println(correct.Render(fmt.Sprintf("%s : %s", text, input)))
			fmt.Println("Great that was the right answer")
			printExample(word.Example)
			fmt.Println("Press enter to continue")

			scanner.Scan()

			clearScreen()
		}
	}
}

// GuessTheWrongWords starts the second loop of words, it goes on until all the words
// have been guessed, the score is no longer updated
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

		if word.Translation != input || len(args) == 0 {
			clearScreen()

			word.WrongCounter++
			wrongWords[text] = word

			fmt.Println(wrong.Render(fmt.Sprintf("%s : %s", text, input)))
			fmt.Println("Expected: ", correctAnswerStyle.Render(word.Translation))
			printExample(word.Example)
			fmt.Println("Press enter to continue")

			scanner.Scan()

			clearScreen()
			continue
		}

		if word.Translation == input {
			clearScreen()

			word.WrongCounter++
			word.Guessed = true

			fmt.Println(correct.Render(fmt.Sprintf("%s : %s", text, input)))
			fmt.Println("Great that was the right answer")
			printExample(word.Example)
			fmt.Println("Press enter to continue")

			scanner.Scan()

			delete(wordMap, text)

			clearScreen()
		}
	}
}

// NewGame starts the guessing word game and resets score and counter at the end
func NewGame(iterations int) {
	scanner := bufio.NewScanner(os.Stdin)
	wordMap := PickRandomWords(iterations)

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

	registerStats(int32(score))

	Reset()

	fmt.Println("Press enter to see the report")
	scanner.Scan()
	GenerateReport(wordMap)
}

func Reset() {
	score = 0
	counter = 0
	exit = false
	wrongWords = make(map[string]*Word)
}
