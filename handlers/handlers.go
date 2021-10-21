package handlers

import (
	"GitHub/goland-twitter/middlew"
	"GitHub/goland-twitter/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func Handlers() {
	router := mux.NewRouter()
	router.HandleFunc("/signup", middlew.CheckBD(routes.SignUp)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckBD(routes.Login)).Methods("POST")
	router.HandleFunc("/profile", middlew.CheckBD(middlew.CheckJWT(routes.SeeProfile))).Methods("GET")
	router.HandleFunc("/update/profile", middlew.CheckBD(middlew.CheckJWT(routes.UpdateProfile))).Methods("PUT")
	router.HandleFunc("/posts", middlew.CheckBD(middlew.CheckJWT(routes.MakePost))).Methods("POST")
	router.HandleFunc("/posts/all", middlew.CheckBD(middlew.CheckJWT(routes.ListTweets))).Methods("GET")
	router.HandleFunc("/posts", middlew.CheckBD(middlew.CheckJWT(routes.TweetsByUser))).Methods("GET")
	router.HandleFunc("/posts", middlew.CheckBD(middlew.CheckJWT(routes.DeleteTweet))).Methods("DELETE")
	router.HandleFunc("/profile/avatar", middlew.CheckBD(middlew.CheckJWT(routes.UploadAvatar))).Methods("POST")
	router.HandleFunc("/profile/banner", middlew.CheckBD(middlew.CheckJWT(routes.UploadBanner))).Methods("POST")
	router.HandleFunc("/profile/avatar", middlew.CheckBD(middlew.CheckJWT(routes.RetrieveAvatar))).Methods("GET")
	router.HandleFunc("/profile/banner", middlew.CheckBD(middlew.CheckJWT(routes.RetrieveBanner))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
