package main

import (
	"bufio"
	"fmt"
	"os"
)

type command struct {
	text     string
	desc     string
	callback func()
}

func getCommands() map[string]command {
	return map[string]command{
		"help": {
			text:     "help",
			desc:     "List of available commands",
			callback: callbackHelp,
		},
		"quit": {
			text:     "quit",
			desc:     "Exit the program",
			callback: callbackExit,
		},
		"play": {
			text:     "play",
			desc:     "Start the flashcards game",
			callback: callbackPlay,
		},
		"new": {
			text:     "new",
			desc:     "Add a new word to the collection",
			callback: callbackNew,
		},
	}
}

func Repl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("What do you want to do?")
		fmt.Printf("> ")
		scanner.Scan()
		input := scanner.Text()
		commands := getCommands()

		command, ok := commands[input]
		if !ok {
			callbackDefault()
			continue
		}
		command.callback()
	}

}
