package poker_test

import (
	"bytes"
	"strings"
	"testing"

	poker "github.com/whiny-nil/learn-go-with-tests/23-build-an-application"
)

func TestCLI(t *testing.T) {
	t.Run("it runs a game and records the winner", func(t *testing.T) {
		in := strings.NewReader("7\nCleo\n")
		stdout := &bytes.Buffer{}
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 7)
		assertGameFinishedWith(t, game, "Cleo")
	})

	t.Run("it runs another game and records that winner", func(t *testing.T) {
		in := strings.NewReader("2\nChris wins\n")
		stdout := &bytes.Buffer{}
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 2)
		assertGameFinishedWith(t, game, "Chris")
	})

	t.Run("it prints an error when a non-numeric value is entered and does not start the game", func(t *testing.T) {
		in := strings.NewReader("X\n")
		stdout := &bytes.Buffer{}
		game := &GameSpy{}
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadInputMsg)
		assertGameNotStarted(t, game)
	})
}

func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()

	want := strings.Join(messages, "")
	got := stdout.String()

	if got != want {
		t.Errorf("got %q sent to stdout but wanted %q", got, want)
	}
}

func assertGameNotStarted(t testing.TB, game *GameSpy) {
	t.Helper()

	if game.StartCalled {
		t.Error("game should not have been started")
	}
}

func assertGameStartedWith(t testing.TB, game *GameSpy, want int) {
	t.Helper()

	if game.StartedWith != want {
		t.Errorf("wanted Start called with %d but got %d", want, game.StartedWith)
	}
}

func assertGameFinishedWith(t testing.TB, game *GameSpy, want string) {
	t.Helper()

	if game.FinishedWith != want {
		t.Errorf("got %s, want %s", game.FinishedWith, want)
	}

}

type GameSpy struct {
	StartCalled  bool
	StartedWith  int
	FinishedWith string
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartCalled = true
	g.StartedWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}
