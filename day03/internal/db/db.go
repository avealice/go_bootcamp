package db

import (
	"bytes"
	"context"
	"day03/internal/types"
	"encoding/json"
	"log"
)

func (es *ElasticsearchStore) GetPlaces(limit int, offset int) ([]types.Place, int, error) {
	if offset >= 13640 {
		return nil, -1, nil
	}

	query := map[string]interface{}{
		"from": offset,
		"size": limit,
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, 0, err
	}

	res, err := es.client.Search(
		es.client.Search.WithContext(context.Background()),
		es.client.Search.WithIndex(es.indexName),
		es.client.Search.WithBody(&buf),
		es.client.Search.WithTrackTotalHits(true),
		es.client.Search.WithPretty(),
	)

	if err != nil {
		log.Fatalf("Elasticsearch Search() API ERROR:", err)
	}
	if res.IsError() {
		log.Fatalf("Elasticsearch Search() API ERROR:", res)
	}

	defer res.Body.Close()

	var response map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, 0, err
	}

	hits := response["hits"].(map[string]interface{})["hits"].([]interface{})

	totalValue := int(response["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))

	var places []types.Place

	for _, hit := range hits {
		hitMap := hit.(map[string]interface{})
		source := hitMap["_source"].(map[string]interface{})
		place := types.Place{
			ID:      int(source["id"].(float64)),
			Name:    source["name"].(string),
			Address: source["address"].(string),
			Phone:   source["phone"].(string),
			Location: types.Location{
				Lon: source["location"].(map[string]interface{})["lon"].(float64),
				Lat: source["location"].(map[string]interface{})["lat"].(float64),
			},
		}
		places = append(places, place)
	}

	return places, totalValue, nil
}

func (es *ElasticsearchStore) GetRecommendations(lat, lon float64) ([]types.Place, error) {
	query := map[string]interface{}{
		"size": 3,
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
		"sort": []map[string]interface{}{
			{
				"_geo_distance": map[string]interface{}{
					"location": map[string]interface{}{
						"lat": lat,
						"lon": lon,
					},
					"order":           "asc",
					"unit":            "km",
					"mode":            "min",
					"distance_type":   "arc",
					"ignore_unmapped": true,
				},
			},
		},
	}

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	res, err := es.client.Search(
		es.client.Search.WithContext(context.Background()),
		es.client.Search.WithIndex(es.indexName),
		es.client.Search.WithBody(&buf),
		es.client.Search.WithTrackTotalHits(true),
		es.client.Search.WithPretty(),
	)

	if err != nil {
		log.Fatalf("Elasticsearch Search() API ERROR:", err)
	}
	if res.IsError() {
		log.Fatalf("Elasticsearch Search() API ERROR:", res)
	}

	defer res.Body.Close()

	var response map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	hits := response["hits"].(map[string]interface{})["hits"].([]interface{})

	var places []types.Place

	for _, hit := range hits {
		hitMap := hit.(map[string]interface{})
		source := hitMap["_source"].(map[string]interface{})
		place := types.Place{
			ID:      int(source["id"].(float64)),
			Name:    source["name"].(string),
			Address: source["address"].(string),
			Phone:   source["phone"].(string),
			Location: types.Location{
				Lon: source["location"].(map[string]interface{})["lon"].(float64),
				Lat: source["location"].(map[string]interface{})["lat"].(float64),
			},
		}
		places = append(places, place)
	}

	return places, nil
}
