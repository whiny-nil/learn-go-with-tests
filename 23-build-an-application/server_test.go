package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

// server_test.go
func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		[]string{},
	}
	server := &PlayerServer{&store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, 200)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, 200)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns a 404 if the player is missing", func(t *testing.T) {
		request := newGetScoreRequest("Zippy")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, 404)
	})
}

// server_test.go
func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		[]string{},
	}
	server := &PlayerServer{&store}

	t.Run("it records a win on POST", func(t *testing.T) {
		player := "Bob"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, http.StatusAccepted)
		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin, wanted %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("got %q calls, wanted %q", store.winCalls[0], player)
		}
	})
}

func newGetScoreRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func newPostWinRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func assertResponseStatus(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got status %d, want status %d", got, want)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got body %q, want body %q", got, want)
	}
}
