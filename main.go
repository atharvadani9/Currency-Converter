package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	"Currency-Converter/internal/app"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
		log.Println("Make sure to set APP_ID environment variable manually")
	}
	fmt.Println("🌍 Currency Converter TUI")

	application := app.NewApplication()
	if err := application.Run(); err != nil {
		log.Printf("Application error: %v", err)
		os.Exit(1)
	}
	fmt.Println("\nThanks for using Currency Converter!")
}
