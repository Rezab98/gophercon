package main

import (
	"log"
	"os"

	"github.com/Rezab98/gophercon/internal/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		log.Printf("Error executing command: %v", err)
		os.Exit(1)
	}
}
