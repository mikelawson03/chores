package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mikelawson03/chores/internal/api"
	"github.com/mikelawson03/chores/internal/app"
	"github.com/mikelawson03/chores/internal/domain"
	"github.com/mikelawson03/chores/internal/events"
	"github.com/mikelawson03/chores/internal/middleware"
	"github.com/mikelawson03/chores/internal/store"
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
	devUsername := os.Getenv("DEV_USERNAME")
	devPassword := os.Getenv("DEV_PASSWORD")
	JWTSigninSecret := os.Getenv("JWT_SIGNIN_SECRET")

	dbConn := dbConnect(dbPath)
	dbMigrate(dbConn)

	eventBus := events.NewBus()
	config := app.NewConfig(JWTSigninSecret, devPassword, devUsername)
	store := store.NewStore(dbConn)
	app := app.NewApp(store, config, eventBus)
	go eventBus.Listen()

	func() {
		ctx := context.Background()
		user1, err := store.GetHouseholdUserByID(ctx, "840a61a5-4a12-485c-9bf4-d373aa3074ff", domain.DefaultHouseholdID)
		if err != nil {
			fmt.Println(err)
		}

		user2, err := store.GetHouseholdUserByID(ctx, "ae792ee8-c635-44f2-bcad-9bd2142d1c4d", domain.DefaultHouseholdID)
		if err != nil {
			fmt.Println(err)
		}

		id1, events1 := eventBus.Subscribe(user1, events.SubscriberTypeClient)
		id2, events2 := eventBus.Subscribe(user2, events.SubscriberTypeClient)
		go func() {
			for event := range events1 {
				fmt.Println("subscriber1:", id1, event)
			}
		}()

		go func() {
			for event := range events2 {
				fmt.Println("subscriber2:", id2, event)
			}
		}()
	}()

	cfg := api.NewApiConfig(app)

	m := http.NewServeMux()
	cfg.RegisterRoutes(m)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: middleware.Cors(m),
	}

	log.Printf("Server open and listening on port: %s", port)
	log.Fatal(srv.ListenAndServe())

}
