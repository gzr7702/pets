package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gzr7702/pets/models"

	"github.com/labstack/echo"
)

// GetOwners endpoint
func GetOwners(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetOwners(db))
	}
}

// PutOwner endpoint
func PutOwner(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// create a new owner
		var owner models.Owner

		// map incoming JSON to new owner
		c.Bind(&owner)

		//Add a new owner using model
		id, err := models.PutOwner(db, owner.Name)

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

// DeleteOwner endpoint
func DeleteOwner(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		// use model to delete owner
		_, err := models.DeleteOwner(db, id)
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
