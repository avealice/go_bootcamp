package dbcompare

import (
	"encoding/xml"
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
