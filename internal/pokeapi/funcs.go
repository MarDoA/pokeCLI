package pokeapi

import (
	"encoding/json"

	"github.com/MarDoA/pokeCLI/internal/pokecache"
)

func GetLocationAreaList(cfg *Config, s string, cach *pokecache.Cache) ([]string, error) {
	url := ""
	switch s {
	case "next":
		url = cfg.Next
	case "prev":
		url = cfg.Previous
	}
	if url == "" {
		url = baseURL + "/location-area/"
	}
	var data []byte
	if dat, ok := cach.Get(url); ok {
		data = dat
	} else {
		dat, err := getData(url)
		if err != nil {
			return nil, err
		}
		data = dat
		cach.Add(url, data)
	}
	var laresp locationAreaListResponse
	err := json.Unmarshal(data, &laresp)
	if err != nil {
		return nil, err
	}

	var areas []string
	for _, d := range laresp.Results {
		areas = append(areas, d.Name)
	}
	cfg.Next = laresp.Next
	cfg.Previous = laresp.Previous
	return areas, nil
}

func GetPokemonInArea(area string, cach *pokecache.Cache) (string, []string, error) {
	url := baseURL + "/location-area/" + area
	var data []byte
	if dat, ok := cach.Get(url); ok {
		data = dat
	} else {
		dat, err := getData(url)
		if err != nil {
			return "", nil, err
		}
		data = dat
		cach.Add(url, data)
	}
	var laresp locationAreaResponse
	err := json.Unmarshal(data, &laresp)
	if err != nil {
		return "", nil, err
	}
	var pokemons []string
	for _, pokemon := range laresp.PokemonEncounters {
		pokemons = append(pokemons, pokemon.Pokemon.Name)
	}
	return laresp.Name, pokemons, nil
}

func GetPokemon(pokemon string, cach *pokecache.Cache) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon
	var data []byte
	if dat, ok := cach.Get(url); ok {
		data = dat
	} else {
		dat, err := getData(url)
		if err != nil {
			return Pokemon{}, err
		}
		data = dat
		cach.Add(url, data)
	}
	var pokeresp Pokemon
	err := json.Unmarshal(data, &pokeresp)
	if err != nil {
		return Pokemon{}, err
	}
	return pokeresp, nil

}
