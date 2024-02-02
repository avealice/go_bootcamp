package web

import (
	"day03/internal/db"
	"day03/internal/types"
	"html/template"
	"net/http"
	"strconv"
)

type PageData struct {
	Total    int
	Places   []types.Place
	HasPrev  bool
	PrevPage int
	HasNext  bool
	NextPage int
	LastPage int
}

func Handler(es *db.ElasticsearchStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pageStr := r.URL.Query().Get("page")
		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			http.Error(w, "Invalid 'page' value: '"+pageStr+"'", http.StatusBadRequest)
			return
		}

		limit := 10
		offset := (page - 1) * limit

		places, total, err := es.GetPlaces(limit, offset)

		if total < 0 {
			http.Error(w, "Invalid 'page' value: '"+pageStr+"'", http.StatusBadRequest)
			return
		}

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		lastPage := total / limit
		hasPrev := page > 1
		hasNext := page < lastPage

		if page > lastPage {
			http.Error(w, "Invalid 'page' value: '"+pageStr+"'", http.StatusBadRequest)
			return
		}

		data := PageData{
			Total:    total,
			Places:   places,
			HasPrev:  hasPrev,
			PrevPage: page - 1,
			HasNext:  hasNext,
			NextPage: page + 1,
			LastPage: lastPage,
		}

		tmpl, err := template.ParseFiles("web/template/template.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}
