package routes

import (
	"GitHub/goland-twitter/bd"
	"GitHub/goland-twitter/models"
	"encoding/json"
	"log"
	"net/http"
)

func RetrieveRelation(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("id")) == 0 {
		http.Error(w, "UserRelationId is required", 400)
		return
	}

	rel := models.Relationship{
		UserRelationId: r.URL.Query().Get("id"),
		UserID:         UserID,
	}

	status, err := bd.RetrieveRelation(rel)
	if err != nil {
		log.Print(w, "The record cannot be retrieved -- "+err.Error())
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.RelationResponseRetrieve{
		Status: status,
	})
}
