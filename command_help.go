package main

import (
	"fmt"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	commands := getCommands()
	var commandOrder = []string{"help", "map", "mapb", "explore", "catch", "inspect", "pokedex", "exit"}

	for _, name := range commandOrder {
		cmd := commands[name]
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()

	return nil
}
