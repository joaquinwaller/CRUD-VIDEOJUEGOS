package main

import (
	"CRUD-VIDEOJUEGOS/internal/service"
	"CRUD-VIDEOJUEGOS/internal/store"
	"CRUD-VIDEOJUEGOS/internal/transport"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./videogames.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	q := `
	CREATE TABLE IF NOT EXISTS videogames(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		online BOOLEAN NOT NULL
	);`

	if _, err = db.Exec(q); err != nil {
		log.Fatal(err)
	}

	videogamesStore := store.New(db)
	videogamesService := service.New(videogamesStore)
	videogamesHandler := transport.New(videogamesService)

	http.HandleFunc("/videogames", videogamesHandler.HandleVideogames)
	http.HandleFunc("/videogames/", videogamesHandler.HandleVideogamesByID)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
