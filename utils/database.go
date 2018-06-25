package utils

import "database/sql"

func InitDB(filepath string) *sql.DB {
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

func Migrate(db *sql.DB) {
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
