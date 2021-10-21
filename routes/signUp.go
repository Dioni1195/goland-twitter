package routes

import (
	"GitHub/goland-twitter/bd"
	"GitHub/goland-twitter/models"
	"encoding/json"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error with request data "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "The password`s length must be greater than six", 400)
		return
	}

	_, found, _ := bd.CheckUserAlreadyExist(t.Email)
	if found == true {
		http.Error(w, "The user already exists", 400)
		return
	}

	_, status, err := bd.CreateUser(t)
	if err != nil {
		http.Error(w, "Error while saving the record "+err.Error(), 500)
		return
	}

	if status == false {
		http.Error(w, "User doesnt create", 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
