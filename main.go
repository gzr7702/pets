package main

import (
	"database/sql"
	"github.com/gzr7702/pets/handlers"
	"github.com/labstack/echo"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbPath := "/home/rob/go/src/github.com/gzr7702/pets/storage.db"
	indexPath := "/home/rob/go/src/github.com/gzr7702/pets/public/index.html"
	db := initDB(dbPath)
	migrate(db)

	// new instance of echo
	e := echo.New()

	// routes
	e.File("/", indexPath)
	e.GET("/pets", handlers.GetPets(db))
	e.PUT("/pets", handlers.PutPet(db))
	e.DELETE("/pets/:id", handlers.DeletePet(db))

	// start the web server
	e.Start(":8000")
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	// check for db errors and exit
	if err != nil {
		panic(err)
	}

	// exit if we don't get a db connection
	if db == nil {
		panic("db nil")
	}

	return db

}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS pets(
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR NOT NULL,
			animal_type VARCHAR NOT NULL
		);
	`

	// in go, underbar is used to assign an unused var
	print("execute\n")
	_, err := db.Exec(sql)

	// exit if something goes wrong w/sql statement
	if err != nil {
		panic(err)
	}

}
