package poker_test

import (
	"bytes"
	"strings"
	"testing"

	poker "github.com/whiny-nil/learn-go-with-tests/23-build-an-application"
)

func TestCLI(t *testing.T) {
	t.Run("it runs a game and records the winner", func(t *testing.T) {
		in := strings.NewReader("7\nCleo wins\n")
		stdout := &bytes.Buffer{}
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		poker.AssertGameStartedWith(t, game, 7)
		poker.AssertGameFinishedWith(t, game, "Cleo")
	})

	t.Run("it runs another game and records that winner", func(t *testing.T) {
		in := strings.NewReader("2\nChris wins\n")
		stdout := &bytes.Buffer{}
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		poker.AssertGameStartedWith(t, game, 2)
		poker.AssertGameFinishedWith(t, game, "Chris")
	})

	t.Run("it prints an error when a non-numeric value is entered and does not start the game", func(t *testing.T) {
		in := strings.NewReader("X\n")
		stdout := &bytes.Buffer{}
		game := &poker.GameSpy{}
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadInputMsg)
		poker.AssertGameNotStarted(t, game)
	})

	t.Run("it does not finish the game until '<name> wins' is entered", func(t *testing.T) {
		in := strings.NewReader("2\nCleo\nCleo wins\n")
		stdout := &bytes.Buffer{}
		game := &poker.GameSpy{}
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.InstructionText)
		poker.AssertGameStartedWith(t, game, 2)
		poker.AssertGameFinishedWith(t, game, "Cleo")
	})

	t.Run("it does not finish the game when the string 'quit' is entered", func(t *testing.T) {
		in := strings.NewReader("2\nquit\n")
		stdout := &bytes.Buffer{}
		game := &poker.GameSpy{}
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.GoodbyeMsg)
		poker.AssertGameStartedWith(t, game, 2)
		poker.AssertGameNotFinished(t, game)
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
