package api

import (
	"day03/internal/db"
	"encoding/json"
	"net/http"
	"strconv"
)

func writeJSONResponse(w http.ResponseWriter, response map[string]interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func PlacesHandler(es *db.ElasticsearchStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pageStr := r.URL.Query().Get("page")
		page, err := strconv.Atoi(pageStr)

		if err != nil || page < 1 {
			response := map[string]interface{}{"error": "Invalid 'page' value: '" + pageStr + "'"}
			writeJSONResponse(w, response, http.StatusBadRequest)
			return
		}

		limit := 10
		offset := (page - 1) * limit

		places, total, err := es.GetPlaces(limit, offset)
		if err != nil {
			writeJSONResponse(w, map[string]interface{}{"error": "Internal Server Error"}, http.StatusInternalServerError)
			return
		}

		lastPage := total / limit

		if page > lastPage {
			response := map[string]interface{}{"error": "Invalid 'page' value: '" + pageStr + "'"}
			writeJSONResponse(w, response, http.StatusBadRequest)
			return
		}

		response := map[string]interface{}{
			"name":      "Places",
			"total":     total,
			"places":    places,
			"prev_page": page - 1,
			"next_page": page + 1,
			"last_page": lastPage,
		}

		writeJSONResponse(w, response, http.StatusOK)
	}
}

func RecommendHandler(es *db.ElasticsearchStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		latStr := r.URL.Query().Get("lat")
		lonStr := r.URL.Query().Get("lon")

		lat, err := strconv.ParseFloat(latStr, 64)
		if err != nil {
			http.Error(w, "Invalid 'lat' value", http.StatusBadRequest)
			return
		}

		lon, err := strconv.ParseFloat(lonStr, 64)
		if err != nil {
			http.Error(w, "Invalid 'lon' value", http.StatusBadRequest)
			return
		}

		recommendations, err := es.GetRecommendations(lat, lon)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"name":   "Recommendation",
			"places": recommendations,
		}

		writeJSONResponse(w, response, http.StatusOK)
	}
}

func HandleGetToken(w http.ResponseWriter, r *http.Request) {
	token, err := GenerateToken()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"token": token}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
