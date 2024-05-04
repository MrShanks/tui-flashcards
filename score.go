package main

import (
	"fmt"
	"os"
	"strconv"
)

func UpdateBestScore(filename string, score int, wordMap map[string]string) error {
	scoreStr := strconv.Itoa(score)
	err := os.WriteFile(filename, []byte(scoreStr), 0644)
	if err != nil {
		return err
	}
	fmt.Println("Congrats this is the best you have done so far")
	fmt.Println(correct.Render(fmt.Sprintf("%d/%d", score, len(wordMap))))
	fmt.Printf("Don't forget to come back tomorrow!\n")
	return nil
}

func ReadBestScore(filename string) (int, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	bestScore, err := strconv.Atoi(string(content))
	if err != nil {
		return 0, err
	}
	return bestScore, nil

}
