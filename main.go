package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	// Internal Packages
	"github.com/AviralxD/the-homely-casino/internal/database"
	"github.com/AviralxD/the-homely-casino/routes"

	// External Packages
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	err := godotenv.Load(".env")

	portName := os.Getenv("PORT")
	if portName == "" {
		log.Fatal("Port not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL not found in the environment")
	}

	Conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to db:", err)
	}

	dbConn := database.New(Conn)

	apiCfg := apiConfig{
		DB: dbConn,
	}

	router := http.NewServeMux()
	routes.Route(router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portName,
	}

	log.Printf("server starting on port:%v", portName)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port:", portName)
}
