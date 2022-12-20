package poker

import (
	"io"
	"time"
)

type TexasHoldem struct {
	playerStore PlayerStore
	alerter     BlindAlerter
}

func NewTexasHoldem(playerStore PlayerStore, alerter BlindAlerter) *TexasHoldem {
	return &TexasHoldem{
		playerStore: playerStore,
		alerter:     alerter,
	}
}

func (game *TexasHoldem) Start(numberOfPlayers int, alertDestination io.Writer) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		game.alerter.ScheduleAlertAt(blindTime, blind, alertDestination)
		blindTime = blindTime + blindIncrement
	}
}

func (game *TexasHoldem) Finish(playerName string) {
	game.playerStore.RecordWin(playerName)
}
