package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gzr7702/pets/models"

	"github.com/labstack/echo"
)

type H map[string]interface{}

// GetPets endpoint
func GetPets(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetPets(db))
	}
}

// CreatePet endpoint
func CreatePet(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// create a new pet
		var pet models.Pet

		// map incoming JSON to new pet
		c.Bind(&pet)

		//Add a new pet using model
		id, err := models.CreatePet(db, pet.Name, pet.Type, pet.OwnerID)

		// return JSON if successful
		// LastInertId available error is a bug in postgres driver. Ignore it
		if err == nil || err.Error() == "no LastInsertId available" {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})

		} else {
			return err
		}

	}
}

// UpdatePetName endpoint
func UpdatePetName(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		name := c.Param("name")

		// use model to update pet
		_, err := models.UpdatePetName(db, id, name)
		// return JSON response if successful
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"updated": id,
			})
		} else {
			return err
		}
	}
}

// UpdatePetOwner endpoint
func UpdatePetOwner(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		ownerID, _ := strconv.Atoi(c.Param("ownerid"))

		// use model to update pet
		_, err := models.UpdatePetOwner(db, id, ownerID)
		// return JSON response if successful
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"updated": id,
			})
		} else {
			return err
		}
	}
}

// DeletePet endpoint
func DeletePet(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		// use model to delete pet
		_, err := models.DeletePet(db, id)
		// return JSON response if successful
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		} else {
			return err
		}
	}
}
