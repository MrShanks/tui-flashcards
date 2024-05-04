package main

import (
	"bufio"
	"fmt"
	"os"
)

var score, counter int
var wrongWords = make(map[string]string)

func startUp(scanner *bufio.Scanner) {
	clearScreen()
	fmt.Println("Hi Welcome back to your language training")
	fmt.Println("Ready?")
	scanner.Scan()
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	words, err := LoadWords()
	if err != nil {
		fmt.Println(err)
	}
	wordMap := pickRandomWords(words, 10)

	bestScore, err := ReadBestScore("bestScore.txt")
	if err != nil {
		fmt.Println(err)
	}

	startUp(scanner)
	translateWords(wordMap, scanner, false)

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
}
