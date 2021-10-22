package routes

import (
	"GitHub/goland-twitter/bd"
	"encoding/json"
	"net/http"
	"strconv"
)

func ListFollowersTweets(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Page is required and must be greater than Zero "+err.Error(), http.StatusBadRequest)
		return
	}

	response, status := bd.ListFollowersTweets(UserID, page)
	if !status {
		http.Error(w, "Error getting tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
