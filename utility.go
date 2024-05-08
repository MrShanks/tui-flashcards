package main

import (
	"fmt"
	"os"
	"os/exec"
)

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printExample(example string) {
	fmt.Println("Here is an example:")
	fmt.Println(exampleStyle.Render(example))
}
