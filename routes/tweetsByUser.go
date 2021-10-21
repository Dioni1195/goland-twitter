package routes

import (
	"GitHub/goland-twitter/bd"
	"encoding/json"
	"net/http"
	"strconv"
)

func TweetsByUser(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	if len(userId) == 0 {
		http.Error(w, "User Id is required", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) == 0 {
		http.Error(w, "Page is required", http.StatusBadRequest)
		return
	}

	page, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)
	if err != nil {
		http.Error(w, "Data error "+err.Error(), http.StatusBadRequest)
		return
	}
	if page == 0 {
		http.Error(w, "Page must be greater than Zero", http.StatusBadRequest)
		return
	}

	tweets, err := bd.RetrieveTweetsByUser(userId, page)
	if err != nil {
		http.Error(w, "Error retrieving tweets "+err.Error(), 500)
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweets)
}
