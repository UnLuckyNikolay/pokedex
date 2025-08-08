package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Pokedex!")
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		command := reader.Text()
		words := cleanInput(command)
		if len(words) == 0 {
			continue
		}

		cmd, ok := commandRegistry[words[0]]
		if !ok {
			fmt.Printf("Command '%s' not found!\n", words[0])
			continue
		}
		cmd.callback()
	}
}
