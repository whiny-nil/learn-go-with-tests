package main

import (
	"log"
	"os"

	poker "github.com/whiny-nil/learn-go-with-tests/23-build-an-application"
)

const dbFileName = "game.db.json"

func main() {
	store, closeDB, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	alerter := poker.BlindAlerterFunc(poker.StdOutAlerter)
	game := poker.NewTexasHoldem(store, alerter)

	if err != nil {
		log.Fatal(err)
	}
	defer closeDB()

	cli := poker.NewCLI(os.Stdin, os.Stdout, game)
	cli.PrintWelcome()
	cli.PlayPoker()
}
