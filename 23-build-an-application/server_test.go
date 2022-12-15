package poker

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
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
	server := NewPlayerServer(&store)

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
	server := NewPlayerServer(&store)

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
		server := NewPlayerServer(&store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertResponseStatus(t, response.Code, http.StatusOK)
		AssertResponseContentType(t, response.Result().Header.Get("content-type"), "application/json")
		got := getLeagueFromResponse(t, response.Body)
		AssertLeague(t, got, wantedLeague)
	})
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
