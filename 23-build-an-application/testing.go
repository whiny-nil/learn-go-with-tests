package poker

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

type ScheduledAlert struct {
	At     time.Duration
	Amount int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

type SpyBlindAlerter struct {
	Alerts []ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.Alerts = append(s.Alerts, ScheduledAlert{duration, amount})
}

func AssertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	if len(store.winCalls) != 1 {
		t.Errorf("got %d calls to RecordWin, wanted %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("got %q as winner, wanted %q", store.winCalls[0], winner)
	}
}

func AssertLeague(t *testing.T, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func AssertResponseStatus(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got status %d, want status %d", got, want)
	}
}

func AssertResponseContentType(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got status %q, want status %q", got, want)
	}
}

func AssertResponseBody(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got body %q, want body %q", got, want)
	}
}

func AssertScoreEquals(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got score %d, want score %d", got, want)
	}
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error, but got one, %v", err)
	}
}

func AssertScheduledAlert(t *testing.T, got, want ScheduledAlert) {
	t.Helper()

	if got.Amount != want.Amount {
		t.Errorf("got amount %d, want %d", got.Amount, want.Amount)
	}

	if got.At != want.At {
		t.Errorf("got scheduled time of %v, want %v", got.At, want.At)
	}

}