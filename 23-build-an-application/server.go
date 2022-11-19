package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	player := strings.TrimPrefix(req.URL.Path, "/players/")
	switch req.Method {
	case http.MethodPost:
		p.processWin(res, player)
	case http.MethodGet:
		p.showScore(res, player)
	}
}

func (p *PlayerServer) processWin(res http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	res.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(res http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		res.WriteHeader(http.StatusNotFound)
	} else {
		fmt.Fprint(res, score)
	}
}

func GetPlayerScore(player string) int {
	switch player {
	case "Pepper":
		return 20
	case "Floyd":
		return 10
	default:
		return 0
	}
}

func RecordWin(player string) {}
