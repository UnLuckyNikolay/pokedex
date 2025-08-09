# Pokedex

A CLI Pokemon exploration tool built in Go. Integrates with [PokeAPI](https://pokeapi.co), features a REPL loop and caching.

## Install and Run

1. Install [Go](https://go.dev/doc/install) 1.24.5 or higher:

2. Clone the repository:

    ```bash
	git clone https://github.com/UnLuckyNikolay/pokedex
    cd pokedex
	```

3. Build and run:

	```bash 
	go build 
	./pokedex
	```

	or run without building:

	```bash
	go run .
	```

## Available Commands

* `help` - Prints the list of available commands.
* `map` - Displays 20 next locations.
* `mapb` - Displays 20 previous locations.
* `explore <id/location>` - Moves you into the specified location. Leave empty to reexplore current location.
* `catch <pokemon>` - Tries to catch the specified pokemon. You need to be in the same location as them.
* `inspect <pokemon>` - Inspect the pokemon that was previously caught.
* `pokedex` - Shows the list of the caught pokemon.
* `exit` - Exits the Pokedex.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.