package main

import (
	"github.com/MrShanks/tui-flashcards/game"
	_ "github.com/lib/pq"
)

func main() {
	game.Repl()
}
