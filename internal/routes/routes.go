package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/NNNiv/pokedex/db"
	"github.com/NNNiv/pokedex/internal/templates/pages"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func HomeHandler(c echo.Context) error {
	return Render(c, http.StatusOK, pages.Home())
}

func SurpriseMeHandler(c echo.Context) error {
	return Render(c, http.StatusOK, pages.Surprise())
}

func WhoHandler(c echo.Context) error {
	return Render(c, http.StatusOK, pages.Who())
}

func IDKHandler(c echo.Context) error {
	return Render(c, http.StatusOK, pages.Idk())
}

func SearchHandler(c echo.Context) error {
	pokemon := c.FormValue("searchForm")
	var pokemonData db.Pokemon
	var err error

	if isNumeric(pokemon) {
		if len(pokemon) == 1 {
			pokemon = "00" + pokemon
		} else if len(pokemon) == 2 {
			pokemon = "0" + pokemon
		}
		pokemonData, err = db.GetPokemonByID(pokemon)
	} else {
		pokemonData, err = db.GetPokemonByName(pokemon)
	}

	if err != nil {
		return nil
	}
	fmt.Println(pokemonData.Type1)
	fmt.Println(*pokemonData.Type2)
	return nil
}

func isNumeric(str string) bool {
	if _, err := strconv.Atoi(str); err == nil {
		return true
	}
	return false
}
