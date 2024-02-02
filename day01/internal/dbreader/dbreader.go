package dbreader

import (
	"encoding/xml"
	"path/filepath"
)

type DBReader interface {
	ReadRecipes(filePath string) (*DataBase, error)
	PrintRecipes(db DataBase) error
}

type Recipe struct {
	Name        string       `json:"name" xml:"name"`
	StoveTime   string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

type Ingredient struct {
	ItemName  string `json:"ingredient_name" xml:"itemname"`
	ItemCount string `json:"ingredient_count" xml:"itemcount"`
	ItemUnit  string `json:"ingredient_unit,omitempty" xml:"itemunit,omitempty"`
}

type DataBase struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Recipes []Recipe `json:"cake" xml:"cake"`
}

func NewReader(filePath string) DBReader {
	switch ext := getFileExtension(filePath); ext {
	case ".json":
		return &JSONReader{}
	case ".xml":
		return &XMLReader{}
	default:
		return nil
	}
}

func getFileExtension(filePath string) string {
	return filepath.Ext(filePath)
}
