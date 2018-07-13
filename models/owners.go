package models

import (
	"database/sql"
)

type Owner struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type OwnerCollection struct {
	Owners []Owner `json:"owners"`
}

// GetOwners gets all pets from database
func GetOwners(db *sql.DB) OwnerCollection {
	sql := "SELECT * FROM owners"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	result := OwnerCollection{}
	for rows.Next() {
		owner := Owner{}
		err2 := rows.Scan(&owner.ID, &owner.Name)

		if err2 != nil {
			panic(err2)
		}
		result.Owners = append(result.Owners, owner)
	}
	return result
}

// PutOwner puts an owner into database
func PutOwner(db *sql.DB, name string) (int64, error) {
	sql := "INSERT INTO owners(name) VALUES($1)"

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err2 := stmt.Exec(name)

	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

// UpdateOwnerName updates a single pet from the database
func UpdateOwnerName(db *sql.DB, id int, name string) (int64, error) {
	sql := "UPDATE owners SET name = $2 WHERE owner_id = $1"

	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err2 := stmt.Exec(id, name)

	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}

// DeleteOwner deletes a single pet from the database
func DeleteOwner(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM owners WHERE owner_id = $1"

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
