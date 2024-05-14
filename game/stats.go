package game

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/MrShanks/tui-flashcards/internal/database"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func registerStats(score int32) {
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
	stat := database.CreateStatParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Score:     score,
	}
	_, err = queries.CreateStat(ctx, stat)
	if err != nil {
		log.Fatal("Can't create the stat: ", err)
	}
}
