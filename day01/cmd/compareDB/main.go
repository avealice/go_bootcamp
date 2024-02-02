package main

import (
	"flag"
	"fmt"
	"s21/internal/dbcompare"
)

func main() {
	oldFilePath := flag.String("old", "", "Path to the original database XML/JSON file")
	newFilePath := flag.String("new", "", "Path to the stolen database XML/JSON file")
	flag.Parse()

	if *oldFilePath == "" || *newFilePath == "" {
		fmt.Println("Usage: ./compareDB --old original_database.xml --new stolen_database.json")
		return
	}

	var oldReader, newReader = dbcompare.NewReader(*oldFilePath), dbcompare.NewReader(*newFilePath)
	if oldReader == nil || newReader == nil {
		fmt.Printf("Error: Invalid file extension\n")
		return
	}

	oldDB, err := oldReader.ReadRecipes(*oldFilePath)

	if err != nil {
		fmt.Printf("Error reading recipes: %v", err)
		return
	}

	newDB, err := newReader.ReadRecipes(*newFilePath)

	if err != nil {
		fmt.Printf("Error reading recipes: %v", err)
		return
	}

	dbcompare.CompareDataBases(*oldDB, *newDB)
	return
}
