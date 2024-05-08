package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

var wordFile = "new_words.txt"

type Word struct {
	text         string
	translation  string
	example      string
	wrongCounter int
}

func NewWord(text, translation, example string) *Word {
	return &Word{
		text:        text,
		translation: translation,
		example:     example,
	}
}

func AddWordToFile() error {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Write the new German word")
	fmt.Printf("> ")
	scanner.Scan()
	germanWord := scanner.Text()

	fmt.Println("Write the English translation")
	fmt.Printf("> ")
	scanner.Scan()
	translation := scanner.Text()

	fmt.Println("Type a sentence with that word")
	fmt.Printf("> ")
	scanner.Scan()
	sentence := scanner.Text()

	newEntry := fmt.Sprintf("%s,%s,%s\n", germanWord, translation, sentence)

	file, err := os.OpenFile(wordFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	_, err = file.WriteString(newEntry)
	if err != nil {
		return err
	}

	return nil
}

func pickRandomWords(words []*Word, n int) map[string]*Word {
	wordMap := make(map[string]*Word)
	copyWords := make([]*Word, len(words))
	copy(copyWords, words)
	rand.Shuffle(len(copyWords), func(i, j int) {
		copyWords[i], copyWords[j] = copyWords[j], copyWords[i]
	})
	for _, word := range copyWords[:n] {
		wordMap[word.text] = word
	}
	return wordMap
}

// LoadWords reads words from the words.txt file and creates a slice of words
// that are stored in memory
func LoadWords() ([]*Word, error) {
	file, err := os.Open(wordFile)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var words []*Word

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		if len(parts) != 3 {
			fmt.Println("Invalid line:", scanner.Text())
			continue
		}

		word := NewWord(
			strings.TrimSpace(parts[1]),
			strings.TrimSpace(parts[0]),
			strings.TrimSpace(parts[2]))
		words = append(words, word)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return words, nil
}

func ListWordsFromFile() {
	words, err := LoadWords()
	if err != nil {
		fmt.Printf("Could Load words from file: %s\n", err)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "German", "English", "Example"})
	for i, word := range words {
		i++
		t.AppendRow(table.Row{i, word.translation, word.text, word.example})
		t.AppendSeparator()
	}
	t.SetStyle(table.StyleRounded)
	t.Render()
}
