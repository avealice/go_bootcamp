package dbreader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type XMLReader struct{}

func (r *XMLReader) ReadRecipes(filePath string) (*DataBase, error) {
	xmlData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var db *DataBase
	err = xml.Unmarshal(xmlData, &db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (r *XMLReader) PrintRecipes(db DataBase) error {
	jsonResult, err := json.MarshalIndent(db, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(string(jsonResult))
	return nil
}
