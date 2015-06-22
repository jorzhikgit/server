package main

import (
	"errors"
)

type Game struct {
	Id             int
	Name           string
	AvailableItems []Item // items available when creating a board
	Players        []Player
	Theme          string // not a choice, just a suggestion from the Host
}

type GameRepository interface {
	Save(game Game)
	FindById(id int) Game
	FindHost() Player
}

// Creates a new game
func NewGame(name string, theme string, host Player) Game {
	players := make([]Player, 0, 8)
	players = append(players, host)

	return Game{
		Id:             2,
		Name:           name,
		Theme:          theme,
		Players:        players,
		AvailableItems: make([]Item, 0),
	}
}

// Add manages adding new items to the list of available items in the game
func (game *Game) AddToAvailable(item Item) (int, error) {
	// do not allow duplicate items
	for _, existingItem := range game.AvailableItems {
		if existingItem.Value == item.Value {
			return len(game.AvailableItems), errors.New("Game can not have duplicate items available")
		}
	}

	game.AvailableItems = append(game.AvailableItems, item)
	return len(game.AvailableItems), nil
}
