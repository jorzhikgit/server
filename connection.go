package main

import (
	"github.com/boyle-bingo/server/events"
	"github.com/gorilla/websocket"
)

// Connection is the interface that manages the physical reading of
// inputs from the user and the writing of events back to the user
type Connection interface {

	// Read messages from connection interface
	Read() error

	// Write messages to connection
	Write(Event events.Event) error

	// Close this connection
	Close() error
}

// WsConnection is is the Connection implmentation for websockets
type WsConnection struct {

	// Websocket connection from gorilla
	ws *websocket.Conn

	// outbound messages
	send chan []byte
}

// NewWsConnection creates a new connection struct from a Websocket connection
// as well as a channel for messages that are meant to be passed back to the
// user
func NewWsConnection(wsCon *websocket.Conn) *WsConnection {
	return &WsConnection{
		ws:   wsCon,
		send: make(chan []byte, 0),
	}
}

// Read messages received from the user of this connection
func (w *WsConnection) Read() error {
	return nil
}

// Write events back to the user. Events are marshalled into JSON
func (w *WsConnection) Write(Event events.Event) error {
	return nil
}

// Close this websocket connection
func (w *WsConnection) Close() error {
	return nil
}
