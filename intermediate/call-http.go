package intermediate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func SampleHttpConsume() {
	response,err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil{
		log.Fatal(err)
	}

	//print as string
	fmt.Println(string(responseData))

	//convert
	var pokemonResponse Response
	json.Unmarshal(responseData,&pokemonResponse)
	fmt.Println("----")
	fmt.Println(pokemonResponse.Name," * ",len(pokemonResponse.Pokemon))

	for i:=0 ; i<len(pokemonResponse.Pokemon) ; i++  {
		fmt.Println(pokemonResponse.Pokemon[i].Species.Name)
	}
}

type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
}