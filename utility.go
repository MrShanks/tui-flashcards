package main

import (
	"fmt"
	"os"
	"os/exec"
)

// clearScreen cleans up the screen on a mac
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// printExample prints the example in Blue
func printExample(example string) {
	fmt.Println("Here is an example:")
	fmt.Println(exampleStyle.Render(example))
}
