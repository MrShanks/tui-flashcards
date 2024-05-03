package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

var normal = lipgloss.NewStyle().
	Border(lipgloss.ThickBorder(), true, false).
	BorderForeground(lipgloss.Color("#3C3C3C")).
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#3C3C3C")).
	PaddingTop(1).
	PaddingBottom(1).
	Align(lipgloss.Center).
	Width(22)

var wrong = lipgloss.NewStyle().
	Border(lipgloss.ThickBorder(), true, false).
	BorderForeground(lipgloss.Color("#3C3C3C")).
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#E7625F")).
	PaddingTop(1).
	PaddingBottom(1).
	Align(lipgloss.Center).
	Width(22)

var correctAnswerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00EE00"))

var correct = lipgloss.NewStyle().
	Border(lipgloss.ThickBorder(), true, false).
	BorderForeground(lipgloss.Color("#3C3C3C")).
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#8BCD50")).
	PaddingTop(1).
	PaddingBottom(1).
	Align(lipgloss.Center).
	Width(22)

var score int

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

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

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	scanner := bufio.NewScanner(os.Stdin)
	bestScore, err := ReadBestScore("bestScore.txt")

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

	if score > bestScore {
		UpdateBestScore("bestScore.txt", score)
		fmt.Println("Congrats this is the best you have done so far")
		fmt.Println(normal.Render(fmt.Sprintf("%d/%d", score, len(words))))
		fmt.Printf("Don't forget to come back tomorrow!\n")
	} else {
		fmt.Printf("Great your final Score is: \n")
		fmt.Println(normal.Render(fmt.Sprintf("%d/%d", score, len(words))))
		fmt.Printf("Don't forget to come back tomorrow!\n")

	}
}
