package main

import (
	"database/sql"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	api "github.com/mikelawson03/chores/internal/api"
	"github.com/mikelawson03/chores/internal/middleware"
	"github.com/pressly/goose"
)

func dbConnect(dbPath string) *sql.DB {
	goose.SetDialect("sqlite3")

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("connected to db")

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("db pinged")

	return db
}

func dbMigrate(dbConn *sql.DB) {
	before, _ := goose.GetDBVersion(dbConn)

	goose.SetLogger(log.New(io.Discard, "", 0))

	err := goose.Up(dbConn, "./sql/schema")
	if err != nil {
		log.Fatal(err)
	}

	after, _ := goose.GetDBVersion(dbConn)

	if after > before {
		log.Printf("Database migrated from version %d to %d", before, after)
	} else {
		log.Printf("Database already at latest version (%d)", after)
	}
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	dbPath := os.Getenv("DB_PATH")

	dbConn := dbConnect(dbPath)
	dbMigrate(dbConn)

	cfg := api.NewApiConfig(dbConn)

	m := http.NewServeMux()
	cfg.RegisterRoutes(m)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: middleware.Cors(m),
	}

	log.Printf("Server open and listening on port: %s", port)
	srv.ListenAndServe()
}
