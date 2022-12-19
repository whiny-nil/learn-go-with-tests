package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

func (cli *CLI) PlayPoker() {
	cli.printPrompt()

	numberOfPlayersInput := cli.readLine()
	numberOfPlayers, err := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))

	if err != nil {
		cli.printBadInput()
		return
	}

	cli.game.Start(numberOfPlayers)
	userInput := cli.readLine()
	cli.game.Finish(extractWinner(userInput))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func (cli *CLI) PrintInstructions() {
	cli.printSection()
	cli.printInstructionText()
	cli.printSection()
}

func (cli *CLI) printInstructionText() {
	fmt.Fprint(cli.out, "Let's play poker!\n")
	fmt.Fprint(cli.out, "\n")
	fmt.Fprint(cli.out, "Type '<Name> wins' to record a win,\n")
	fmt.Fprint(cli.out, "or 'quit' to exit without recording a win\n")
}

func (cli *CLI) printSection() {
	fmt.Fprint(cli.out, "\n")
	fmt.Fprint(cli.out, "------------------------------------------\n")
	fmt.Fprint(cli.out, "\n")
}

const PlayerPrompt = "Please enter the number of players: "

func (cli *CLI) printPrompt() {
	fmt.Fprint(cli.out, PlayerPrompt)
}

const BadInputMsg = "You must enter a number for the number of players\n"

func (cli *CLI) printBadInput() {
	fmt.Fprint(cli.out, BadInputMsg)
}
