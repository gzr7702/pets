package main

import (
	"github.com/gzr7702/pets/handlers"
	"github.com/labstack/echo"
	"path"

	"github.com/gzr7702/pets/utils"

	_ "github.com/lib/pq"
)

func main() {
	appRoot := "/home/rob/go/src/github.com/gzr7702/pets"

	indexPath := path.Join(appRoot, "assets/index.html")
	jsPath := path.Join(appRoot, "assets/js/app.js")

	db := utils.InitDB()
	utils.Migrate(db)

	// new instance of echo
	e := echo.New()

	// static files
	e.File("/", indexPath)
	e.File("/app.js", jsPath)

	// pet routes
	e.GET("/pets", handlers.GetPets(db))
	e.POST("/pets", handlers.CreatePet(db))
	e.PUT("/pets/:id/name/:name", handlers.UpdatePetName(db))
	e.PUT("/pets/:id/owner/:ownerid", handlers.UpdatePetOwner(db))
	e.DELETE("/pets/:id", handlers.DeletePet(db))

	// owner routes
	e.GET("/owners", handlers.GetOwners(db))
	e.POST("/owners", handlers.PutOwner(db))
	e.PUT("/owners/:id/name/:name", handlers.UpdateOwnerName(db))
	e.DELETE("/owners/:id", handlers.DeleteOwner(db))

	// start the web server
	e.Start(":8000")
}
