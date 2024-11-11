package main

import (
	project "OlxScraper"
	"OlxScraper/internal/api/router"
	sqlcDb "OlxScraper/internal/db"
	"OlxScraper/internal/repository"
	"OlxScraper/internal/service"
	"database/sql"
	"github.com/caarlos0/env/v6"
	"github.com/pressly/goose/v3"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"

	_ "embed"
	"log"
	"os"
)

type Config struct {
	SqliteFile string `env:"SQLITE_FILE,required"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}

	db, err := setupDatabase(cfg.SqliteFile)
	if err != nil {
		log.Fatal(err)
	}

	queries := sqlcDb.New(db)

	repo := repository.New(queries)
	svc := service.New(repo)
	r := router.New(svc)

	port := os.Getenv("PORT")

	r.Logger.Fatal(r.Start(":" + port))
}

func setupDatabase(sqliteFile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", sqliteFile)
	if err != nil {
		return nil, err
	}

	if err := runMigrations(db); err != nil {
		return nil, err
	}

	return db, nil
}

func runMigrations(db *sql.DB) error {
	goose.SetBaseFS(project.EmbedMigrations)

	if err := goose.SetDialect("sqlite"); err != nil {
		panic(err)
	}
	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}

	log.Println("Migrations applied successfully.")
	return nil
}
