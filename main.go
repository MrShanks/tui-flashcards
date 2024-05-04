package main

import (
	"bufio"
	"fmt"
	"os"
)

var score, counter int
var wrongWords = make(map[string]string)

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
}
