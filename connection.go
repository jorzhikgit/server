package main

import (
	"errors"

	"github.com/gorilla/websocket"
)

const NULL_GAME int = 0

type Connection interface {

	// Read messages from connection interface
	Read() error

	// Write messages to connection
	Write() error

	CloseChannel() error

	JoinGame(gameId int) error

	LeaveGame()

}

type WsConnection struct {
	ws *websocket.Conn

	// outbound messages
	send chan []byte

	// the game the player is in
	currentGame int
}

func NewWsConnection(wsCon *websocket.Conn) WsConnection {
	return WsConnection {
		ws: wsCon,
		send: make(chan []byte, 0),
		currentGame: NULL_GAME,
	}
}

func (w *WsConnection) Read() error {
	return nil
}

func (w *WsConnection) Write() error {
	return nil
}

func (w *WsConnection) CloseChannel() {
}

func (w *WsConnection) JoinGame(gameId int) error {
	if w.currentGame != NULL_GAME {
		return errors.New("Currently in a game")
	}

	w.currentGame = gameId

	return nil
}

func (w *WsConnection) LeaveGame() {
	w.currentGame = NULL_GAME
}
