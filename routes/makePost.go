package routes

import (
	"GitHub/goland-twitter/bd"
	"GitHub/goland-twitter/models"
	"encoding/json"
	"net/http"
	"time"
)

func MakePost(w http.ResponseWriter, r *http.Request) {
	var createTweet models.CreateTweet
	var tweet models.Tweet

	err := json.NewDecoder(r.Body).Decode(&tweet)
	if err != nil {
		http.Error(w, "Error with request data "+err.Error(), 400)
		return
	}

	if len(tweet.Message) == 0 {
		http.Error(w, "Message is required", 400)
		return
	}

	createTweet = models.CreateTweet{
		UserID:  UserID,
		Message: tweet.Message,
		Date:    time.Now(),
	}

	_, status, err := bd.CreateTweet(createTweet)
	if err != nil {
		http.Error(w, "Error while saving the record "+err.Error(), 500)
		return
	}

	if !status {
		http.Error(w, "Tweet doesnt create", 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
