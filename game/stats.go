package game

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/MrShanks/tui-flashcards/internal/database"
	"github.com/google/uuid"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/joho/godotenv"
)

const (
	YYYYMMDD = "2006-01-02"
	hhmm     = "15:05"
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

func GetAllStats() {
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
	stats, err := queries.GetStats(ctx)
	if err != nil {
		log.Fatal("Can't retrieve scores from database: ", err)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Date", "Time", "Score"})
	for i, stat := range stats {
		t.AppendRow(table.Row{i, stat.CreatedAt.Format(YYYYMMDD), stat.CreatedAt.Format(hhmm), stat.Score})
		t.AppendSeparator()
	}
	t.SetStyle(table.StyleRounded)
	t.Render()
}
