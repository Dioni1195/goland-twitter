package routes

import (
	"GitHub/goland-twitter/bd"
	"encoding/json"
	"net/http"
)

func SeeProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	user, err := bd.FindUser(ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
