package controller

import (
	"encoding/json"
	"lucytech/internal/models"
	"lucytech/internal/service"
	"net/http"
)

func GetHTMLContextMetadata(w http.ResponseWriter, r *http.Request) {
	var requestDto = models.HtmlParseRequestDto{}

	if err := json.NewDecoder(r.Body).Decode(&requestDto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseDto := service.ParseHTML(requestDto.Url)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	err := json.NewEncoder(w).Encode(responseDto)

	if err != nil {
		return
	}
}
