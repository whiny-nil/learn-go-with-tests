package poker

import "time"

type Game struct {
	playerStore PlayerStore
	alerter     BlindAlerter
}

func NewGame(playerStore PlayerStore, alerter BlindAlerter) *Game {
	return &Game{
		playerStore: playerStore,
		alerter:     alerter,
	}
}

func (game *Game) Start(numberOfPlayers int) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		game.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}

func (game *Game) Finish(playerName string) {
	game.playerStore.RecordWin(playerName)
}
