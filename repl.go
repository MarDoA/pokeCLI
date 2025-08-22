package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func start() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		command, exist := commands[input[0]]
		if exist {
			err := command.callback()
			if err != nil {
				println(err)
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
