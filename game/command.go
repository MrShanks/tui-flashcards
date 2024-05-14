package game

type command struct {
	text     string
	desc     string
	callback func()
}

// getCommands returns a map of commands
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
		"list": {
			text:     "list",
			desc:     "List all words in the database",
			callback: callbackList,
		},
	}
}
