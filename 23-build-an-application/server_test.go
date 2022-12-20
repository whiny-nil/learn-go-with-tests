package poker

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

var (
	dummyGame = &GameSpy{}
)

// server_test.go
func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		[]string{},
		nil,
	}
	server := mustMakePlayerServer(t, &store, dummyGame)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertResponseStatus(t, response.Code, 200)
		AssertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertResponseStatus(t, response.Code, 200)
		AssertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns a 404 if the player is missing", func(t *testing.T) {
		request := newGetScoreRequest("Zippy")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertResponseStatus(t, response.Code, 404)
	})
}

// server_test.go
func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		[]string{},
		nil,
	}
	server := mustMakePlayerServer(t, &store, dummyGame)

	t.Run("it records a win on POST", func(t *testing.T) {
		player := "Bob"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertResponseStatus(t, response.Code, http.StatusAccepted)
		AssertPlayerWin(t, &store, player)
	})
}

func TestLeague(t *testing.T) {
	t.Run("it returns some JSON /league", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Marc", 14},
		}
		store := StubPlayerStore{nil, nil, wantedLeague}
		server := mustMakePlayerServer(t, &store, dummyGame)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertResponseStatus(t, response.Code, http.StatusOK)
		AssertResponseContentType(t, response.Result().Header.Get("content-type"), "application/json")
		got := getLeagueFromResponse(t, response.Body)
		AssertLeague(t, got, wantedLeague)
	})
}

func TestGame(t *testing.T) {
	t.Run("GET /game returns 200", func(t *testing.T) {
		server := mustMakePlayerServer(t, &StubPlayerStore{}, dummyGame)

		request := newGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertResponseStatus(t, response.Code, http.StatusOK)
	})

	t.Run("start a game with 3 players and declare Ruth the winner", func(t *testing.T) {
		game := &GameSpy{}
		winner := "Ruth"
		server := httptest.NewServer(mustMakePlayerServer(t, &StubPlayerStore{}, game))
		ws := mustDialWS(t, "ws"+strings.TrimPrefix(server.URL, "http")+"/ws")

		defer server.Close()
		defer ws.Close()

		writeWSMessage(t, ws, "3")
		writeWSMessage(t, ws, winner)

		time.Sleep(10 * time.Millisecond)
		AssertGameStartedWith(t, game, 3)
		AssertGameFinishedWith(t, game, winner)
	})
}

func newGameRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/game", nil)
	return request
}

func newLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func getLeagueFromResponse(t *testing.T, body io.Reader) (league []Player) {
	t.Helper()
	league, err := NewLeague(body)

	if err != nil {
		t.Fatal(err)
	}

	return league
}

func newGetScoreRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func newPostWinRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func mustMakePlayerServer(t *testing.T, store PlayerStore, game *GameSpy) *PlayerServer {
	server, err := NewPlayerServer(store, game)
	if err != nil {
		t.Fatal("problem creating player server", err)
	}
	return server
}

func mustDialWS(t *testing.T, url string) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		t.Fatalf("could not open a ws connection on %s %v", url, err)
	}

	return ws
}

func writeWSMessage(t testing.TB, conn *websocket.Conn, message string) {
	t.Helper()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("could not send message over ws connection %v", err)
	}
}
