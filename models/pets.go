package models

import (
	"database/sql"
)

type Pet struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	OwnerID int    `json:"ownerId"`
}

type PetCollection struct {
	// TODO: change this to pets from items =============================================
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
		err2 := rows.Scan(&pet.ID, &pet.OwnerID, &pet.Name, &pet.Type)

		if err2 != nil {
			panic(err2)
		}
		result.Pets = append(result.Pets, pet)
	}
	return result
}

// PutPet puts a pet into database
func PutPet(db *sql.DB, name string, animalType string, ownerID int) (int64, error) {
	sql := "INSERT INTO pets(name, animal_type, owner_id) VALUES($1, $2, $3)"

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	// if owner doesn't exist, make the store the owner
	if ownerID == 0 {
		ownerID = 1
	}

	result, err2 := stmt.Exec(name, animalType, ownerID)

	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

// DeletePet deletes a single pet from the database
func DeletePet(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM pets WHERE pet_id = $1"

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
