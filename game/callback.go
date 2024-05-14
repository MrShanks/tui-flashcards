package game

import (
	"fmt"
	"os"
)

// callbackExit terminates the program immediately
func callbackExit() {
	fmt.Println("Ok bye, see you!")
	os.Exit(0)
}

// callbackHelp lists all the available commands on the screen
func callbackHelp() {
	fmt.Println("These are the available commands")
	for text, cmd := range getCommands() {
		fmt.Printf("- %s: %s\n", text, cmd.desc)
	}
}

// callbackPlay starts a new game
func callbackPlay() {
	NewGame(iterations)
}

// callbackNew enables the user to add a new word in the dictionary
func callbackNew() {
	AddWordToFile()
}

// callbackList lists all the words from the dictionary
func callbackList() {
	ListWordsFromFile()
}

// callbackDefault prints the default error message
func callbackDefault() {
	fmt.Println("Invalid command, type help if you are stuck!")
}
