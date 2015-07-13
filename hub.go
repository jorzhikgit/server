package main

import (
	"encoding/json"
)

const NULL_GAME int = 0

type Hub struct {
	// A connection mapped to a GameId
	gameConnections map[int][]*User

	broadcast chan Event

	// Register connection into hub
	register chan *User

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
		broadcast:       make(chan Event),
		register:        make(chan *User),
		unregister:      make(chan *User),
		gameConnections: make(map[int][]*User),
	}
}

// Listen for events and manage connections
func (h *Hub) Run() {
	for {
		select {
		case u := <-h.register:
			h.gameConnections[NULL_GAME] = append(h.gameConnections[NULL_GAME], u)
		//case u := <-h.unregister:
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

// Removes game id when a player leaves a game
func (h *Hub) LeaveGame(Conn Connection) {
	//h.connections[Conn] = 0
}
