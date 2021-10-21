package routes

import (
	"GitHub/goland-twitter/bd"
	"GitHub/goland-twitter/jwt"
	"GitHub/goland-twitter/models"
	"encoding/json"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error with username and/or password"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}

	u, status := bd.TryLogin(t.Email, t.Password)
	if status == false {
		http.Error(w, "Error with username and/or password", 400)
		return
	}

	jwt, err := jwt.GenerateToken(u)
	if err != nil {
		http.Error(w, "Error generating session"+err.Error(), 500)
		return
	}

	response := models.LoginResponse{
		Token: jwt,
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwt,
		Expires: expirationTime,
	})

}
