package routes

import (
	"GitHub/goland-twitter/bd"
	"net/http"
)

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	idTweet := r.URL.Query().Get("id")
	if len(idTweet) == 0 {
		http.Error(w, "Tweet ID is required", http.StatusBadRequest)
		return
	}

	err := bd.DeleteTweet(idTweet, UserID)
	if err != nil {
		http.Error(w, "Error deleting tweet "+err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)

}
