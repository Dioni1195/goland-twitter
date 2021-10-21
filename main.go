package main

import (
	"GitHub/goland-twitter/bd"
	"GitHub/goland-twitter/handlers"
	"log"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Not connection to Mongo")
	}

	handlers.Handlers()
}
