package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Nome    string    `json:"name"`
	Pokemon []pokemon `json:"pokemon_entries"`
}

type pokemon struct {
	Numero  int            `json:"entry_number"`
	Especie PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Nome string `json:"name"`
}

func main() {
	response, err := http.Get("https://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	for _, pokemon := range responseObject.Pokemon {
		fmt.Println(pokemon.Numero, pokemon.Especie.Nome)
	}
}
