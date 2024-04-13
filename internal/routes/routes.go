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

var pokemonData db.Pokemon
var err error
var tries = 3

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

	pokemonData, err = db.GetRandomPokemon()

	if err != nil {
		return nil
	}

	return Render(c, http.StatusOK, pages.Who(pokemonData, tries))
}

func IDKHandler(c echo.Context) error {
	return Render(c, http.StatusOK, pages.Idk())
}

func SearchHandler(c echo.Context) error {
	pokemon := c.FormValue("searchForm")

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

	return Render(c, http.StatusOK, pages.Result(pokemonData))
}

func GuessHandler(c echo.Context) error {
	guess := c.FormValue("guessForm")

	if guess == pokemonData.Name {
		fmt.Println("Correct Guess")
		return Render(c, http.StatusOK, pages.PokemonColorImage(pokemonData))
	} else {
		if tries > 1 {
			tries--
			fmt.Println(tries)
			return Render(c, http.StatusOK, pages.PokemonBWImage(pokemonData, tries))
		} else {
			tries = 3
			return Render(c, http.StatusOK, pages.PokemonColorImage(pokemonData))
		}
	}
}

func isNumeric(str string) bool {
	if _, err := strconv.Atoi(str); err == nil {
		return true
	}
	return false
}

// func DisplayResult(c echo.Context) {
//
// }

// func AdvancedSearchDisplayHandler(c echo.Context) error {
// 	return Render(c, http.StatusOK, components.AdvancedSearch())
// }
