package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"lucytech/internal/models"
	"lucytech/internal/service"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(userID)

	if err != nil {
		return
	}
}

func GetHTMLContextMetadata(w http.ResponseWriter, r *http.Request) {
	var requestDto = models.HtmlParseRequestDto{}

	if err := json.NewDecoder(r.Body).Decode(&requestDto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	title, version := service.ParseHTML(requestDto.Url)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	err := json.NewEncoder(w).Encode(map[string]interface{}{"title": title, "version": version})

	if err != nil {
		return
	}
}
