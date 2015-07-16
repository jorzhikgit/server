package main

import (
	"encoding/json"
)

type Event struct {
	GameId   int
	PlayerId int
	Type     string
	Data     json.RawMessage
}
