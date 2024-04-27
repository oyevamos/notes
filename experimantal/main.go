package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"
	//"your_project/config" // Импортируйте ваш пакет config
	"github.com/oyevamos/notes.git/config"

	_ "github.com/lib/pq"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.DBHost, conf.DBPort, conf.DBUser, conf.DBPassword, conf.DBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	insertStmt := `INSERT INTO notes (user_id, title, content, date_created, date_modified) VALUES ($1, $2, $3, $4, $5)`
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10000000; i++ {
		userID := rand.Intn(999) + 1
		title := fmt.Sprintf("Title %d", rand.Intn(10000000))
		content := fmt.Sprintf("Content %d", rand.Intn(10000000))
		dateCreated := time.Now()
		dateModified := time.Now()

		_, err := db.Exec(insertStmt, userID, title, content, dateCreated, dateModified)
		if err != nil {
			fmt.Println("Error inserting:", err)
			continue
		}

		if i%100000 == 0 {
			fmt.Println(i, "rows inserted")
		}
	}
	fmt.Println("Insertion complete")
}
