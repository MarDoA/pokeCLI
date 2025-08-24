package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/MarDoA/pokeCLI/internal/pokeapi"
	"github.com/MarDoA/pokeCLI/internal/pokecache"
)

func start() {
	sCache := pokecache.NewCache(5 * time.Second)
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	config := &pokeapi.Config{
		PokeDex: map[string]pokeapi.Pokemon{},
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		para := ""
		if len(input) > 1 {
			para = input[1]
		}
		command, exist := commands[input[0]]
		if exist {
			err := command.callback(config, &sCache, para)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
		println("command doesn't exist")
	}
}

func cleanInput(txt string) []string {
	txt = strings.ToLower(txt)
	cleaned := strings.Fields(txt)
	return cleaned
}
