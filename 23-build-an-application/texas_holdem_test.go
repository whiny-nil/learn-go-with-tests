package poker_test

import (
	"testing"
	"time"

	poker "github.com/whiny-nil/learn-go-with-tests/23-build-an-application"
)

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}

// var dummyStdOut = &bytes.Buffer{}

func TestTexasHoldem_Start(t *testing.T) {

	t.Run("it schedules alerts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(dummyPlayerStore, blindAlerter)

		game.Start(5)

		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, want := range cases {
			t.Run(want.String(), func(t *testing.T) {

				if len(blindAlerter.Alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
				}

				got := blindAlerter.Alerts[i]
				poker.AssertScheduledAlert(t, got, want)
			})
		}
	})

	t.Run("it schedules alerts on game start for 7 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(dummyPlayerStore, blindAlerter)

		game.Start(7)

		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		for i, want := range cases {
			t.Run(want.String(), func(t *testing.T) {

				if len(blindAlerter.Alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
				}

				got := blindAlerter.Alerts[i]
				poker.AssertScheduledAlert(t, got, want)
			})
		}
	})
}

func TestTexasHoldem_Finish(t *testing.T) {
	t.Run("record Chris win", func(t *testing.T) {
		playerStore := &poker.StubPlayerStore{}
		game := poker.NewTexasHoldem(playerStore, dummyBlindAlerter)

		game.Finish("Chris")
		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

}
