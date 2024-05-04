package main

import (
	"os"
	"strconv"
)

func UpdateBestScore(filename string, score int) error {
	scoreStr := strconv.Itoa(score)
	err := os.WriteFile(filename, []byte(scoreStr), 0644)
	if err != nil {
		return err
	}
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
