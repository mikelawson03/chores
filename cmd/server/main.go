package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	api "github.com/mikelawson03/chores/internal/api"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	cfg := api.NewApiConfig()

	m := http.NewServeMux()
	cfg.RegisterRoutes(m)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: m,
	}

	fmt.Println("Server open and listening on port:", port)
	srv.ListenAndServe()
}
