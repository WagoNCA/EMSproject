package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectPostgres() {
	connStr := "host=postgres port=5432 user=admin password=adminer dbname=database sslmode=disable"

	var err error

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	for i := 0; i < 10; i++ {
		err = DB.Ping()
		if err == nil {
			log.Println("Connected to database successfully")
			return
		}

		log.Println("Waiting for database...")
		time.Sleep(2 * time.Second)
	}

	log.Fatal("Failed to ping database:", err)
}
