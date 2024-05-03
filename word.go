package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Word struct {
	text        string
	translation string
}

func NewWord(text, translation string) *Word {
	return &Word{
		text:        text,
		translation: translation,
	}
}

func pickRandomWords(words []*Word, n int) []*Word {
	copyWords := make([]*Word, len(words))
	copy(copyWords, words)
	rand.Shuffle(len(copyWords), func(i, j int) {
		copyWords[i], copyWords[j] = copyWords[j], copyWords[i]
	})
	return copyWords[:n]
}

// LoadWords reads words from the words.txt file and creates a slice of words
// that are stored in memory
func LoadWords() ([]*Word, error) {
	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var words []*Word

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		if len(parts) != 2 {
			fmt.Println("Invalid line:", scanner.Text())
			continue
		}

		word := NewWord(
			strings.TrimSpace(parts[1]),
			strings.TrimSpace(parts[0]))
		words = append(words, word)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return words, nil
}
