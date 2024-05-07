package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

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

func translateWords(wordMap map[string][]string, scanner *bufio.Scanner, repeat bool) {

	for text, wordslice := range wordMap {
		translation := wordslice[0]
		example := wordslice[1]
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
			wrongWords[text] = []string{translation, example}
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
