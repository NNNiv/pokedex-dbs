package main

import (
	"github.com/NNNiv/pokedex/db"
	"github.com/NNNiv/pokedex/internal/routes"
	"github.com/labstack/echo/v4"
)

func main() {

	db.DBConnection()

	app := echo.New()

	app.Static("/static", "static")

	app.GET("/", routes.HomeHandler)
	app.GET("/surprise", routes.SurpriseMeHandler)
	app.GET("/who", routes.WhoHandler)
	app.GET("/idk", routes.IDKHandler)

	app.POST("/search", routes.SearchHandler)
	// fmt.Println("http://localhost:4000")
	app.Logger.Fatal(app.Start(":4000"))
}
