package main

import (
	"bufio"
	"fmt"
	"os"
)

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
