package main

import (
	"flag"
	"fmt"
	"s21/internal/dbreader"
)

func main() {
	filePath := flag.String("f", "", "Path to the database file")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Please provide a path to the database file using the -f flag")
		return
	}

	reader := dbreader.NewReader(*filePath)

	if reader == nil {
		fmt.Printf("Error: Invalid file extension\n")
		return
	}

	recipes, err := reader.ReadRecipes(*filePath)
	if err != nil {
		fmt.Printf("Error reading recipes: %v", err)
		return
	}

	err = reader.PrintRecipes(*recipes)
	if err != nil {
		fmt.Printf("Error printing recipes: %v", err)
		return
	}
	return
}
