package game

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/MrShanks/tui-flashcards/internal/database"
	"github.com/joho/godotenv"
)

// UpdateBestScore overrides the bestscore in the bestScore.txt file and renders
// the score on the terminal
func UpdateBestScore(filename string, score int, wordMap map[string]*Word) error {
	scoreStr := strconv.Itoa(score)
	err := os.WriteFile(filename, []byte(scoreStr), 0644)
	if err != nil {
		return err
	}

	fmt.Println("Congrats this is the best you have done so far")
	fmt.Println(correct.Render(fmt.Sprintf("%d/%d", score, len(wordMap))))
	fmt.Printf("Don't forget to come back tomorrow!\n")

	return nil
}

// ReadBestScore returns the bestScore registered until that point
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

func GetMaxScore() {
	ctx := context.Background()

	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("dbURL is empty, please double check the env conf")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to the database: ", err)
	}
	defer conn.Close()

	queries := database.New(conn)
	maxScore, err := queries.GetMaxScore(ctx)
	if err != nil {
		log.Fatal("Can't retrieve max score from database: ", err)
	}
	fmt.Printf("Best Score: %d\n", maxScore)
}
