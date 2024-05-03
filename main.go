package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
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

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	clearScreen()
	fmt.Println("Hi Welcome back to your language training")
	fmt.Println("Ready?")
	time.Sleep(1500 * time.Millisecond)
	words, err := LoadWords()
	words = pickRandomWords(words, 10)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(os.Stdin)

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
	fmt.Printf("Great your final Score is: \n")
	fmt.Println(normal.Render(fmt.Sprintf("%d/%d", score, len(words))))
	fmt.Printf("Don't forget to come back tomorrow!\n")
}
