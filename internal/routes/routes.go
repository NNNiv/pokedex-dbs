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

var (
	pokemonData      db.Pokemon
	favouritePokemon []db.Pokemon
	err              error
	tries            = 3
	types            []string
	hasRenderedList  bool
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
	pokemonData, err = db.GetRandomPokemon()
	if err != nil {
		return nil
	}
	return Render(c, http.StatusOK, pages.Surprise(pokemonData))
}

func SurpriseHandler(c echo.Context) error {
	return Render(c, http.StatusOK, pages.PokemonImage(pokemonData))
}

func WhoHandler(c echo.Context) error {
	pokemonData, err = db.GetRandomPokemon()
	if err != nil {
		return nil
	}

	return Render(c, http.StatusOK, pages.Who(pokemonData, tries))
}

func FavouritesHandler(c echo.Context) error {
	favouritePokemon, err = GetFavouritePokemon()
	if err != nil {
		return nil
	}
	return Render(c, http.StatusOK, pages.Favourites(favouritePokemon))
}

func IDKHandler(c echo.Context) error {
	return Render(c, http.StatusOK, pages.Idk())
}

func SearchHandler(c echo.Context) error {
	pokemon := c.FormValue("searchForm")

	if pokemon == "" {
		return Render(c, http.StatusOK, pages.EmptyForm())
	}

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

	if pokemonData.Id == "" {
		return Render(c, http.StatusOK, pages.PokemonNotFound())
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

func ShowAdvancedSearch(c echo.Context) error {
	return Render(c, http.StatusOK, pages.AdvancedSearch())
}

func AddPokemonAdvancedHandler(c echo.Context) error {
	pkmnType := c.QueryParam("type")

	types = append(types, pkmnType)

	return Render(c, http.StatusOK, pages.SelectedTypeComponent(pkmnType))
}

func AddFavouriteHandler(c echo.Context) error {
	AddFavouritePokemon(pokemonData)

	return nil
}

func AddFavouritePokemon(pokemon db.Pokemon) error {
	_, err := db.ExecuteQuery(fmt.Sprintf("INSERT INTO favourites VALUES (%s)", pokemon.Id))
	if err != nil {
		return nil
	}
	fmt.Println("Inserted Pokemon")

	return nil
}

func DeleteFavouriteHandler(c echo.Context) error {
	DeleteFavouritePokemon(pokemonData)

	return FavouritesHandler(c)
}

func DeleteFavouritePokemon(pokemon db.Pokemon) error {
	_, err := db.ExecuteQuery(fmt.Sprintf("DELETE FROM favourites WHERE id = %s", pokemon.Id))
	if err != nil {
		return nil
	}
	return nil
}

func GetFavouritePokemon() ([]db.Pokemon, error) {
	pokemonIDs := []string{}
	rows, err := db.ExecuteQuery("SELECT * FROM favourites")
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		pokemonIDs = append(pokemonIDs, id)
	}
	for _, id := range pokemonIDs {
		pokemon, err := db.GetPokemonByID(id)
		favouritePokemon = append(favouritePokemon, pokemon)
		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}
	return favouritePokemon, nil

}

func isNumeric(str string) bool {
	if _, err := strconv.Atoi(str); err == nil {
		return true
	}
	return false
}

func RemoveIndex(s []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}
