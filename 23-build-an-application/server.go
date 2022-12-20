package poker

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/gorilla/websocket"
)

type Player struct {
	Name string
	Wins int
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
	template *template.Template
	game     Game
}

type playerServerWS struct {
	*websocket.Conn
}

const htmlTemplatePath = "game.html"

func NewPlayerServer(store PlayerStore, game Game) (*PlayerServer, error) {
	p := new(PlayerServer)

	tmpl, err := template.ParseFiles(htmlTemplatePath)

	if err != nil {
		return nil, fmt.Errorf("problem opening %s, %v", htmlTemplatePath, err)
	}

	p.template = tmpl
	p.game = game
	p.store = store

	router := http.NewServeMux()
	router.Handle("/game", http.HandlerFunc(p.gameHandler))
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	router.Handle("/ws", http.HandlerFunc(p.webSocket))

	p.Handler = router

	return p, nil
}

const jsonContentType = "application/json"

func (p *PlayerServer) gameHandler(res http.ResponseWriter, req *http.Request) {
	p.template.Execute(res, nil)
}

func (p *PlayerServer) leagueHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", jsonContentType)
	json.NewEncoder(res).Encode(p.store.GetLeague())
}

func (p *PlayerServer) playersHandler(res http.ResponseWriter, req *http.Request) {
	player := strings.TrimPrefix(req.URL.Path, "/players/")

	switch req.Method {
	case http.MethodPost:
		p.processWin(res, player)
	case http.MethodGet:
		p.showScore(res, player)
	}
}

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (p *PlayerServer) webSocket(res http.ResponseWriter, req *http.Request) {
	ws := newPlayerServerWS(res, req)

	numberOfPlayersMsg := ws.WaitForMsg()
	numberOfPlayers, _ := strconv.Atoi(string(numberOfPlayersMsg))
	p.game.Start(numberOfPlayers, io.Discard) //todo: don't discard the blinds messages!

	winnerMsg := ws.WaitForMsg()
	p.game.Finish(string(winnerMsg))
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

func newPlayerServerWS(res http.ResponseWriter, req *http.Request) *playerServerWS {
	conn, err := wsUpgrader.Upgrade(res, req, nil)

	if err != nil {
		log.Printf("problem upgrading connection to WebSockets %v\n", err)
	}

	return &playerServerWS{conn}
}

func (w *playerServerWS) WaitForMsg() string {
	_, msg, err := w.ReadMessage()
	if err != nil {
		log.Printf("error reading from websocket %v\n", err)
	}

	return string(msg)
}
