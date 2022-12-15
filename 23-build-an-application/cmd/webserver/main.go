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

	server := poker.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", server))
}
