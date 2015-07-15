package main

import (
	"encoding/json"
)

const NULL_GAME int = 0

type Hub struct {
	// A connection mapped to a GameId
	gameConnections map[int][]*User

	// all currently running games
	runningGames map[int]*GameInteractor

	// connections not in a game
	notInGame map[*User]bool

	// events to send to connections
	broadcast chan Event

	// Register connection into hub
	register chan *User

	// Track a user from not ingame to a game
	joinGame chan struct {
		User *User
		Game *GameInteractor
	}

	// Remove connection from hub
	unregister chan *User
}

type Event struct {
	GameId   int
	PlayerId int
	Type     string
	Data     json.RawMessage
}

func NewHub() Hub {
	return Hub{
		broadcast:  make(chan Event),
		register:   make(chan *User),
		unregister: make(chan *User),
		joinGame: make(chan struct {
			User *User
			Game *GameInteractor
		}),
		gameConnections: make(map[int][]*User),
	}
}

// Listen for events and manage connections
func (h *Hub) Run() {
	for {
		select {
		case u := <-h.register:
			h.gameConnections[NULL_GAME] = append(h.gameConnections[NULL_GAME], u)
			h.notInGame[u] = true
		case j := <-h.joinGame:
			h.JoinGame(j.User, j.Game)
		case ev := <-h.broadcast:
			go func(ev Event) {
				if inGame, ok := h.gameConnections[ev.GameId]; ok {
					for _, u := range inGame {
						u.Connection.Write(ev)
					}
				}
			}(ev)

		}
	}
}

// Change a connections game id if they move to a new game
func (h *Hub) ChangeGame(User *User, NewGameId int, CurrentGameId int) {
	h.gameConnections[CurrentGameId] = nil
	h.gameConnections[NewGameId] = append(h.gameConnections[NewGameId], User)
}

// Attach a user to this game
func (h *Hub) JoinGame(User *User, Game *GameInteractor) {
	h.notInGame[User] = false

	gameId := Game.Game.Id
	if game, ok := h.runningGames[gameId]; ok {
		game.Users = append(game.Users, User)
	} else {
		h.runningGames[gameId] = Game
	}
}
