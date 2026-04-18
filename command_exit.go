package main

import (
	"fmt"
	"os"
)

func commandExit(url *config, params ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
