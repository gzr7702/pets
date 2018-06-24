package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Pet struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type PetCollection struct {
	Pets []Pet `json:"items"`
}

// GetPets gets all pets from database
func GetPets(db *sql.DB) PetCollection {
	sql := "SELECT * FROM pets"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	result := PetCollection{}
	for rows.Next() {
		pet := Pet{}
		err2 := rows.Scan(&pet.ID, &pet.Name, &pet.Type)

		if err2 != nil {
			panic(err2)
		}
		result.Pets = append(result.Pets, pet)
	}
	return result
}

// PutPet puts a pet into database
func PutPet(db *sql.DB, name string, animalType string) (int64, error) {
	sql := "INSERT INTO pets(name, animal_type) VALUES(?, ?)"

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err2 := stmt.Exec(name, animalType)

	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

// DeletePet deletes a single pet from the database
func DeletePet(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM pets WHERE id = ?"

	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err2 := stmt.Exec(id)

	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
