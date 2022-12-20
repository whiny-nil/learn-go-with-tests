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

	game := poker.NewTexasHoldem(store, poker.BlindAlerterFunc(poker.Alerter))

	server, err := poker.NewPlayerServer(store, game)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":5000", server))
}
