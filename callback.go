package main

import (
	"fmt"
	"os"
)

func callbackExit() {
	fmt.Println("Ok bye, see you!")
	os.Exit(0)
}
func callbackHelp() {
	fmt.Println("These are the available commands")
	for text, cmd := range getCommands() {
		fmt.Printf("- %s: %s\n", text, cmd.desc)
	}
}
func callbackPlay() {
	NewGame()
}
func callbackNew() {
	AddWordToFile()
}
func callbackDefault() {
	fmt.Println("Invalid command, type help if you are stuck!")
}
