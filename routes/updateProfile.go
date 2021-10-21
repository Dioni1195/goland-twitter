package routes

import (
	"GitHub/goland-twitter/bd"
	"GitHub/goland-twitter/models"
	"encoding/json"
	"net/http"
)

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error with request data "+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool

	status, err = bd.UpdateUser(t, UserID)
	if err != nil {
		http.Error(w, "Error updating user "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "User cannot be updated", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)

}
