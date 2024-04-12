package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var database *sql.DB

func DBConnection() {
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "pokedex",
		AllowNativePasswords: true,
	}

	var err error
	database, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := database.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("DB Connected")
}

type Pokemon struct {
	Id     string
	Name   string
	Gen    int
	Type1  string
	Type2  *string
	Height float32
	Weight float32
	Hp     int
	Atk    int
	Def    int
	Spatk  int
	Spdef  int
	Spd    int
	Total  int
	Entry  string
}

func GetPokemonByName(name string) (Pokemon, error) {
	var pokemon Pokemon
	var type2Nullable sql.NullString

	dexData, err := database.Query(`
        SELECT d.id, d.name, d.gen, d.type_1, d.type_2, d.height, d.weight,
               s.hp, s.atk, s.def, s.spl_atk, s.spl_def, s.speed, s.total,
               de.entry1 as dex_entry
        FROM dex AS d
        LEFT JOIN stats AS s ON d.id = s.id
        LEFT JOIN dex_entries AS de ON d.id = de.id
        WHERE d.name = ?;
    `, name)

	if err != nil {
		return pokemon, fmt.Errorf("getPokemonByName %q: %v", name, err)
	}
	defer dexData.Close()

	for dexData.Next() {
		err := dexData.Scan(
			&pokemon.Id, &pokemon.Name, &pokemon.Gen, &pokemon.Type1, &type2Nullable,
			&pokemon.Height, &pokemon.Weight, &pokemon.Hp, &pokemon.Atk, &pokemon.Def,
			&pokemon.Spatk, &pokemon.Spdef, &pokemon.Spd, &pokemon.Total, &pokemon.Entry,
		)
		if err != nil {
			return pokemon, fmt.Errorf("getPokemonByName %q: %v", name, err)
		}
	}
	if type2Nullable.Valid {
		pokemon.Type2 = &type2Nullable.String
	} else {
		pokemon.Type2 = nil
	}
	return pokemon, nil
}

func GetPokemonByID(id string) (Pokemon, error) {
	var pokemon Pokemon
	var type2Nullable sql.NullString

	dexData, err := database.Query(`
        SELECT d.id, d.name, d.gen, d.type_1, d.type_2, d.height, d.weight,
               s.hp, s.atk, s.def, s.spl_atk, s.spl_def, s.speed, s.total,
               de.entry1 as dex_entry
        FROM dex AS d
        LEFT JOIN stats AS s ON d.id = s.id
        LEFT JOIN dex_entries AS de ON d.id = de.id
        WHERE d.id = ?;
    `, id)

	if err != nil {
		return pokemon, fmt.Errorf("getPokemonByID %q: %v", id, err)
	}
	defer dexData.Close()

	for dexData.Next() {
		err := dexData.Scan(
			&pokemon.Id, &pokemon.Name, &pokemon.Gen, &pokemon.Type1, &type2Nullable,
			&pokemon.Height, &pokemon.Weight, &pokemon.Hp, &pokemon.Atk, &pokemon.Def,
			&pokemon.Spatk, &pokemon.Spdef, &pokemon.Spd, &pokemon.Total, &pokemon.Entry,
		)
		if err != nil {
			return pokemon, fmt.Errorf("getPokemonByID %q: %v", id, err)
		}
	}
	if type2Nullable.Valid {
		pokemon.Type2 = &type2Nullable.String
	} else {
		pokemon.Type2 = nil
	}
	return pokemon, nil
}
