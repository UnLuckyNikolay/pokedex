package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		command := reader.Text()
		words := cleanInput(command)
		if len(words) == 0 {
			continue
		}

		fmt.Printf("Your command was: %s\n", words[0])
	}
}
