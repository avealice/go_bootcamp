package main

import (
	"day03/api"
	"day03/internal/db"
	"day03/web"
	"log"
	"net/http"
)

var setAndMap = `{
	"settings": {
		"index": {
			"max_result_window": 20000
		}
	},
	"mappings": {
		"properties": {
		"id": {
			"type": "long"
		},
		"name": {
			"type": "text"
		},
		"address": {
			"type": "text"
		},
		"phone": {
			"type": "text"
		},
		"location": {
			"type": "geo_point"
		}
		}
	}
}`

var indexName = "places"

func main() {
	log.SetFlags(0)

	store, err := db.NewElasticsearchStore(indexName)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	store.CreateIndex(setAndMap)
	store.BulkIndex("../materials/data.csv")

	http.HandleFunc("/", web.Handler(store))
	http.HandleFunc("/api/get_token", api.HandleGetToken)
	http.HandleFunc("/api/places", api.PlacesHandler(store))
	http.HandleFunc("/api/recommend", api.JwtMiddleware(http.HandlerFunc(api.RecommendHandler(store))))
	http.ListenAndServe(":8888", nil)
}
