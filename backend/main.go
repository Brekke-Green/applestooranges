package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	connString := fmt.Sprintf("postgres://%s:%s@localhost:5432/nutritiondb?sslmode=disable", user, password)
	//connString := fmt.Sprint("postgres://postgres:secret@localhost:5432/nutritiondb?sslmode=disable")
	fmt.Printf("%v\n", connString)
	db, err := sql.Open("postgres", connString)
	defer db.Close()

	if err != nil {
		fmt.Print("NOT PING")
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}

