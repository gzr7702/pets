package utils

import "database/sql"

// InitDB initializes database
func InitDB() *sql.DB {

	// TODO: move uid, pwd ot separate file ========================================================================
	db, err := sql.Open("postgres", "user=petsuser password=La0ban dbname=pets sslmode=disable")
	// =============================================================================================================

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
	CREATE TABLE IF NOT EXISTS pets(
			pet_id SERIAL PRIMARY KEY,
			name VARCHAR (50) NOT NULL,
			animal_type VARCHAR (100) NOT NULL
		);
	`

	// in go, underbar is used to assign an unused var
	print("execute\n")
	_, err := db.Exec(sql)

	// exit if something goes wrong w/sql statement
	if err != nil {
		print("Migration error")
		panic(err)
	}

}
