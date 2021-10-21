package routes

import (
	"GitHub/goland-twitter/bd"
	"encoding/json"
	"net/http"
)

func ListTweets(w http.ResponseWriter, r *http.Request) {
	tweets, err := bd.ListTweets()
	if err != nil {
		http.Error(w, "Error retrieving tweets "+err.Error(), 500)
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweets)
}
