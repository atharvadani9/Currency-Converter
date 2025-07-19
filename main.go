package main

import (
	"fmt"
	"log"
	"os"

	"Currency-Converter/internal/app"
)

func main() {
	// Print welcome message
	fmt.Println("🌍 Currency Converter TUI")

	// Create a new application instance
	application := app.NewApplication()

	// Run the TUI application
	if err := application.Run(); err != nil {
		log.Printf("Application error: %v", err)
		os.Exit(1)
	}

	// Exit message
	fmt.Println("\nThanks for using Currency Converter!")
}
