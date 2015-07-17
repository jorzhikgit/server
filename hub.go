package main

import (
	"github.com/boyle-bingo/server/events"
)

type hub interface {
	// Manage interactions with connection
	Run()
}

type Hub struct {
	// A connection mapped to a GameId
	//gameConnections map[int][]*User

	// all currently running games
	runningGames map[int]*GameInteractor

	// connections not in a game
	notInGame map[*User]bool

	// events to send to connections
	broadcast chan events.Event

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

func NewHub() Hub {
	return Hub{
		broadcast:  make(chan events.Event),
		register:   make(chan *User),
		unregister: make(chan *User),
		joinGame: make(chan struct {
			User *User
			Game *GameInteractor
		}),
		notInGame:    make(map[*User]bool),
		runningGames: make(map[int]*GameInteractor),
	}
}

// Listen for events and manage connections
func (h *Hub) Run() {
	for {
		select {
		case u := <-h.register:
			h.notInGame[u] = true
		case j := <-h.joinGame:
			h.JoinGame(j.User, j.Game)
			// probably don't need to broadcast, each game interactor
			// can manage sending to its users
			//case ev := <-h.broadcast:
			//	go func(ev Event) {
			//		if inGame, ok := h.gameConnections[ev.GameId]; ok {
			//			for _, u := range inGame {
			//				u.Connection.Write(ev)
			//			}
			//		}
			//	}(ev)

		}
	}
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
