package routes

import (
	"GitHub/goland-twitter/bd"
	"GitHub/goland-twitter/models"
	"io"
	"net/http"
	"os"
	"strings"
)

func UploadBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	if err != nil {
		http.Error(w, "File is required "+err.Error(), http.StatusBadRequest)
		return
	}
	extension := strings.Split(handler.Filename, ".")[1]
	path := "uploads/banners/" + UserID + "." + extension

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		http.Error(w, "Error creating file"+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error writing file"+err.Error(), http.StatusBadRequest)
		return
	}

	u := models.User{
		Banner: UserID + "." + extension,
	}
	var status bool
	status, err = bd.UpdateUser(u, UserID)
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
