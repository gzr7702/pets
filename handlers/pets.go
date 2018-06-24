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

// PutPet endpoint
func PutPet(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// create a new pet
		var pet models.Pet

		// map incoming JSON to new pet
		c.Bind(&pet)

		//Add a new pet using model
		id, err := models.PutPet(db, pet.Name, pet.Type)

		// return JSON if successful
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
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
