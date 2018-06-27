package utils

import (
	"database/sql"
	"fmt"
)

// InitDB initializes database
func InitDB() *sql.DB {

	credentials := GetInfo()
	serverInfo := fmt.Sprintf("user=%v password=%v dbname=pets sslmode=disable", credentials.User, credentials.Pass)
	db, err := sql.Open("postgres", serverInfo)

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

// Migrate migrates teh database
func Migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS owners(
			owner_id SERIAL PRIMARY KEY,
			name VARCHAR (50) NOT NULL
		);

	CREATE TABLE IF NOT EXISTS pets(
			pet_id SERIAL PRIMARY KEY,
			owner_id INTEGER REFERENCES owners(owner_id),
			name VARCHAR (50) NOT NULL,
			animal_type VARCHAR (100) NOT NULL
		);

	`
	_, err := db.Exec(sql)

	// exit if something goes wrong w/sql statement
	if err != nil {
		print("Migration error")
		panic(err)
	}

}
