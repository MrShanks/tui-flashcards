package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func translateWords(wordMap map[string]string, scanner *bufio.Scanner, repeat bool) {

	for text, translation := range wordMap {
		if !repeat {
			counter++
		}
		clearScreen()
		fmt.Printf("Write the translation %d/%d\t\t\t\tScore: %d\n", counter, len(wordMap), score)
		fmt.Println(normal.Render(text))
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		args := strings.Fields(input)

		if translation != input || len(args) == 0 {
			clearScreen()
			wrongWords[text] = translation
			fmt.Println(wrong.Render(fmt.Sprintf("%s : %s", text, input)))
			fmt.Println("Expected: ", correctAnswerStyle.Render(translation))
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
			fmt.Println("Press enter to continue")
			scanner.Scan()
			clearScreen()
			if repeat {
				delete(wordMap, text)
			}
			continue
		}
	}
}
