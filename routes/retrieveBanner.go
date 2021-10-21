package routes

import (
	"GitHub/goland-twitter/bd"
	"io"
	"net/http"
	"os"
)

func RetrieveBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("userId")
	if len(ID) == 0 {
		http.Error(w, "User Id is required", http.StatusBadRequest)
		return
	}

	user, err := bd.FindUser(ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	openFile, err := os.Open("uploads/banners/" + user.Banner)
	if err != nil {
		http.Error(w, "Image not found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, openFile)
	if err != nil {
		http.Error(w, "Error try later "+err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
}
