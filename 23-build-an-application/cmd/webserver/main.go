package main

import (
	"log"
	"net/http"

	poker "github.com/whiny-nil/learn-go-with-tests/23-build-an-application"
)

const dbFileName = "game.db.json"

func main() {
	store, closeDB, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}

	defer closeDB()

	server, err := poker.NewPlayerServer(store)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":5000", server))
}
