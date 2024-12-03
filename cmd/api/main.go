package main

import (
	"OlxScraper/internal/api/router"
	"OlxScraper/internal/auth"
	sqlcDb "OlxScraper/internal/db"
	"OlxScraper/internal/repository"
	"OlxScraper/internal/service"
	"database/sql"
	"github.com/caarlos0/env/v6"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"

	_ "embed"
	"log"
	"os"
)

type Config struct {
	DbFile    string `env:"DB_FILE,required"`
	JWTSecret string `env:"JWT_SECRET,required"`
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

	db, err := setupDatabase(cfg.DbFile)
	if err != nil {
		log.Fatal(err)
	}

	jwtService := auth.NewJWTService(cfg.JWTSecret)

	queries := sqlcDb.New(db)

	repo := repository.New(queries)
	svc := service.New(repo, jwtService)
	r := router.New(svc, jwtService)

	port := os.Getenv("PORT")

	r.Logger.Fatal(r.Start(":" + port))
}

func setupDatabase(sqliteFile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", sqliteFile)
	if err != nil {
		return nil, err
	}

	return db, nil
}
