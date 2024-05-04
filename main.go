package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var score int
var wrongWords = make(map[string]string)

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	bestScore, err := ReadBestScore("bestScore.txt")
	if err != nil {
		fmt.Println(err)
	}

	clearScreen()

	fmt.Println("Hi Welcome back to your language training")
	fmt.Println("Ready?")
	scanner.Scan()

	words, err := LoadWords()
	words = pickRandomWords(words, 10)
	if err != nil {
		fmt.Println(err)
	}

	for i, word := range words {
		i++
		clearScreen()
		fmt.Printf("Write the translation %d/%d\t\t\t\tScore: %d\n", i, len(words), score)
		fmt.Println(normal.Render(word.text))
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		args := strings.Fields(input)

		if word.translation != input || len(args) == 0 {
			clearScreen()
			wrongWords[word.text] = word.translation
			fmt.Println(wrong.Render(fmt.Sprintf("%s : %s", word.text, input)))
			fmt.Println("Expected: ", correctAnswerStyle.Render(word.translation))
			fmt.Println("Press enter to continue")
			scanner.Scan()
			clearScreen()
			continue
		}

		if word.translation == input {
			clearScreen()
			score++
			fmt.Println(correct.Render(fmt.Sprintf("%s : %s", word.text, input)))
			fmt.Println("Great that was the right answer")
			fmt.Println("Press enter to continue")
			scanner.Scan()
			clearScreen()
			continue
		}
	}

	for {
		if len(wrongWords) == 0 {
			break
		}
		for word, translation := range wrongWords {
			clearScreen()
			fmt.Printf("Write the translation \t\t\t\tScore: %d\n", score)
			fmt.Println(normal.Render(word))
			fmt.Print("> ")
			scanner.Scan()
			input := scanner.Text()
			args := strings.Fields(input)

			if translation != input || len(args) == 0 {
				clearScreen()
				wrongWords[word] = translation
				fmt.Println(wrong.Render(fmt.Sprintf("%s : %s", word, input)))
				fmt.Println("Expected: ", correctAnswerStyle.Render(translation))
				fmt.Println("Press enter to continue")
				scanner.Scan()
				clearScreen()
				continue
			}

			if translation == input {
				clearScreen()
				fmt.Println(correct.Render(fmt.Sprintf("%s : %s", word, input)))
				fmt.Println("Great that was the right answer")
				fmt.Println("Press enter to continue")
				scanner.Scan()
				clearScreen()
				delete(wrongWords, word)
				break
			}
		}
	}
	if score > bestScore {
		UpdateBestScore("bestScore.txt", score)
		fmt.Println("Congrats this is the best you have done so far")
		fmt.Println(correct.Render(fmt.Sprintf("%d/%d", score, len(words))))
		fmt.Printf("Don't forget to come back tomorrow!\n")
	} else {
		fmt.Printf("Great your final Score is: \n")
		fmt.Println(normal.Render(fmt.Sprintf("%d/%d", score, len(words))))
		fmt.Printf("Don't forget to come back tomorrow!\n")

	}
}
