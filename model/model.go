package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	DB_HOST     = "localhost"
	DB_USER     = "postgres"
	DB_PORT     = 5432
	DB_PASSWORD = "123"
	DB_NAME     = "guiwoopark"
)

func DBOpen() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	statement, err := db.Prepare("CREATE TABLE tutorials (id int, tutorial_name text);")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	statement.Exec()
	fmt.Println("Successfully Connected")
}
