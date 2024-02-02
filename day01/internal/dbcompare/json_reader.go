package dbcompare

import (
	"encoding/json"
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
