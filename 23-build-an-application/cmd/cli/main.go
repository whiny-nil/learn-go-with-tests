package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/whiny-nil/learn-go-with-tests/23-build-an-application"
)

const dbFileName = "game.db.json"

func main() {
	store, closeDB, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	alerter := poker.BlindAlerterFunc(poker.StdOutAlerter)
	game := poker.NewGame(store, alerter)

	if err != nil {
		log.Fatal(err)
	}
	defer closeDB()

	fmt.Println("Let's play poker")
	fmt.Println("Type '<Name> wins' to record a win")
	poker.NewCLI(os.Stdin, os.Stdout, game).PlayPoker()
}
