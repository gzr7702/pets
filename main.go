package main

import (
	"github.com/gzr7702/pets/handlers"
	"github.com/labstack/echo"
	"path"

	"github.com/gzr7702/pets/utils"

	_ "github.com/lib/pq"
)

func main() {
	// paths
	appRoot := "/home/rob/go/src/github.com/gzr7702/pets"

	//dbPath := path.Join(appRoot, "storage.db")
	indexPath := path.Join(appRoot, "assets/index.html")
	jsPath := path.Join(appRoot, "assets/js/app.js")

	db := utils.InitDB()
	utils.Migrate(db)

	// new instance of echo
	e := echo.New()

	// routes
	e.File("/", indexPath)
	e.File("/app.js", jsPath)

	e.GET("/pets", handlers.GetPets(db))
	e.PUT("/pets", handlers.PutPet(db))
	e.DELETE("/pets/:id", handlers.DeletePet(db))

	// start the web server
	e.Start(":8000")
}
