package routes

import (
	"GitHub/goland-twitter/bd"
	"GitHub/goland-twitter/models"
	"encoding/json"
	"net/http"
)

func DeleteRelation(w http.ResponseWriter, r *http.Request) {
	var rel models.Relationship

	err := json.NewDecoder(r.Body).Decode(&rel)
	if err != nil {
		http.Error(w, "Error with request data "+err.Error(), 400)
		return
	}
	if len(rel.UserRelationId) == 0 {
		http.Error(w, "UserRelationId is required", 400)
		return
	}

	rel.UserID = UserID

	if err := bd.DeleteRelation(rel); err != nil {
		http.Error(w, "Error while deleting the record "+err.Error(), 500)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
