package dbreader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type JSONReader struct{}

func (r *JSONReader) ReadRecipes(filePath string) (*DataBase, error) {
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var db *DataBase
	err = json.Unmarshal(jsonData, &db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (r *JSONReader) PrintRecipes(db DataBase) error {
	xmlResult, err := xml.MarshalIndent(db, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(string(xmlResult))
	return nil
}
