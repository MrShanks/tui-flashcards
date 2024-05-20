package game

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
	Text         string // English word
	Translation  string // German translation
	Example      string // Real word example
	WrongCounter int    // Number of Attempt
	Guessed      bool   // When guessed is true
}

// NewWord initialize the fields of a newly created word
func NewWord(text, translation, example string) *Word {
	return &Word{
		Text:         text,
		Translation:  translation,
		Example:      example,
		WrongCounter: 0,
		Guessed:      false,
	}
}

// AddWordToFile prompts the user to insert a new word in the dictionary
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

// pickRandomWords shuffles the words in the array and then creates a map of *Words
// with the first n elements
func PickRandomWords(n int) map[string]*Word {
	words, err := LoadWords()
	if err != nil {
		fmt.Println(err)
	}

	wordMap := make(map[string]*Word)
	copyWords := make([]*Word, len(words))
	copy(copyWords, words)
	rand.Shuffle(
		len(copyWords), func(i, j int) {
			copyWords[i], copyWords[j] = copyWords[j], copyWords[i]
		})
	for _, word := range copyWords[:n] {
		wordMap[word.Text] = word
	}
	return wordMap
}

// LoadWords reads words from the new_words.txt file and creates a slice of words
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
		parts := strings.Split(scanner.Text(), ";")
		if len(parts) != 3 {
			fmt.Println("Invalid line:", scanner.Text())
			os.Exit(1)
		}

		word := NewWord(
			strings.TrimSpace(parts[1]),
			strings.TrimSpace(parts[0]),
			strings.TrimSpace(parts[2]))
		words = append(words, word)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
	return words, nil
}

// ListWordsFromFile prints all the available words in the dictionary
// as a table
func ListWordsFromFile() {
	words, err := LoadWords()
	if err != nil {
		fmt.Printf("Couldn't Load words from file: %s\n", err)
		panic(err)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "German", "English", "Example"})
	for i, word := range words {
		i++
		t.AppendRow(table.Row{i, word.Translation, word.Text, word.Example})
		t.AppendSeparator()
	}
	t.SetStyle(table.StyleRounded)
	t.Render()
}
