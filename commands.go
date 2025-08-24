package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"os"

	"github.com/MarDoA/pokeCLI/internal/pokeapi"
	"github.com/MarDoA/pokeCLI/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Config, *pokecache.Cache, string) error
}

func getCommands() map[string]cliCommand {
	Commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display help message",
			callback:    commandHelp,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 location areas",
			callback:    commandMapB,
		},
		"map": {
			name:        "map",
			description: "Display the names of 20 location area, each subsequent call display the next 20",
			callback:    commandMap,
		},
		"explore": {
			name:        "explore",
			description: "Takes a area name or id and displays the name of pokemons in that area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Takes a pokemon name or id and try to catch it",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "takes a pokemon name that was caught and displays its stats",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays the names of your caught pokemons",
			callback:    commandPokedex,
		},
	}
	return Commands
}

func commandExit(c *pokeapi.Config, cach *pokecache.Cache, para string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *pokeapi.Config, cach *pokecache.Cache, para string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, c := range getCommands() {
		fmt.Println(c.name + ": " + c.description)
	}
	return nil
}

func commandMap(c *pokeapi.Config, cach *pokecache.Cache, para string) error {
	areas, err := pokeapi.GetLocationAreaList(c, "next", cach)
	if err != nil {
		return err
	}
	for _, a := range areas {
		fmt.Println(a)
	}
	return nil
}

func commandMapB(c *pokeapi.Config, cach *pokecache.Cache, para string) error {
	if c.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	areas, err := pokeapi.GetLocationAreaList(c, "prev", cach)
	if err != nil {
		return err
	}
	for _, a := range areas {
		fmt.Println(a)
	}
	return nil
}

func commandExplore(c *pokeapi.Config, cach *pokecache.Cache, para string) error {
	if para == "" {
		return errors.New("missing area name or id")
	}
	area, pokemons, err := pokeapi.GetPokemonInArea(para, cach)
	if err != nil {
		return err
	}
	fmt.Println("Exploring " + area + "...")
	fmt.Println("Found Pokemon:")
	for _, p := range pokemons {
		fmt.Println(" - " + p)
	}
	return nil
}

func commandCatch(c *pokeapi.Config, cach *pokecache.Cache, para string) error {
	if para == "" {
		return errors.New("missing pokemon name or id")
	}
	pokemon, err := pokeapi.GetPokemon(para, cach)
	if err != nil {
		return err
	}
	fmt.Println("Throwing a Pokeball at " + pokemon.Name + "...")
	toCatch := pokemon.BaseExp * 100 / 635
	roll := rand.IntN(100)
	if roll < toCatch {
		fmt.Println(pokemon.Name + " escaped!")
		return nil
	}
	fmt.Println(pokemon.Name + " was caught!")
	c.PokeDex[pokemon.Name] = pokemon
	return nil
}

func commandInspect(c *pokeapi.Config, cach *pokecache.Cache, para string) error {
	if para == "" {
		return errors.New("missing pokemon name or id")
	}
	if val, ok := c.PokeDex[para]; ok {
		fmt.Println("Name: " + val.Name)
		fmt.Printf("Height: %d\n", val.Height)
		fmt.Printf("Weight: %d\n", val.Weight)
		fmt.Println("Stats:")
		for _, s := range val.Stats {
			fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range val.Types {
			fmt.Printf("  - %s\n", t.Type.Name)
		}
		return nil
	}
	fmt.Println("you have not caught that pokemon")
	return nil
}

func commandPokedex(c *pokeapi.Config, cach *pokecache.Cache, para string) error {
	if len(c.PokeDex) == 0 {
		fmt.Println("you have no pokemons")
		return nil
	}
	for _, p := range c.PokeDex {
		fmt.Println(" - " + p.Name)
	}
	return nil
}
