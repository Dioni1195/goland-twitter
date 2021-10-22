package routes

import (
	"GitHub/goland-twitter/bd"
	"encoding/json"
	"net/http"
	"strconv"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	param := r.URL.Query().Get("param")

	if len(param) == 0 {
		http.Error(w, "Param is required", http.StatusBadRequest)
		return
	}
	if param != "new" && param != "follow" {
		http.Error(w, "Param must be new or follow", http.StatusBadRequest)
		return
	}

	page, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)
	if err != nil {
		http.Error(w, "Page is required and must be greater than Zero "+err.Error(), http.StatusBadRequest)
		return
	}

	users, status := bd.ListUsers(UserID, page, search, param)
	if !status {
		http.Error(w, "Error getting users", http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}
