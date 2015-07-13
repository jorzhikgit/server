package main

import (
	"github.com/gorilla/websocket"
)

type Connection interface {

	// Read messages from connection interface
	Read() error

	// Write messages to connection
	Write(Event Event) error

	CloseChannel() error
}

type WsConnection struct {
	ws *websocket.Conn

	// outbound messages
	send chan []byte
}

func NewWsConnection(wsCon *websocket.Conn) *WsConnection {
	return &WsConnection{
		ws:   wsCon,
		send: make(chan []byte, 0),
	}
}

func (w *WsConnection) Read() error {
	return nil
}

func (w *WsConnection) Write(Event Event) error {
	return nil
}

func (w *WsConnection) CloseChannel() error {
	return nil
}
