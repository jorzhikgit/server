package main

import (
	"errors"
)

type Game struct {
	Id             int
	Name           string
	AvailableItems []Item // items available when creating a board
	Players        []*Player
	Theme          string // not a choice, just a suggestion from the Host
}

type GameRepository interface {
	Save(game Game) (int, error)
	FindById(id int) (Game, error)

	// Player handling
	FindPlayersByGame(gameId int) ([]Player, error)
}

// NewGame creates a new game structure from a game name, game theme, and
// the host User
func NewGame(name string, theme string, host *Player) Game {
	players := make([]*Player, 0, 8)
	players = append(players, host)

	return Game{
		Id:             2,
		Name:           name,
		Theme:          theme,
		Players:        players,
		AvailableItems: make([]Item, 0),
	}
}

// AddToAvailable manages adding new items to the list of available items in the game
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

// GetHost finds the host in the list of Players
func (game *Game) GetHost() (*Player, error) {
	for _, p := range game.Players {
		if p.IsHost == true {
			return p, nil
		}
	}

	return &Player{}, errors.New("Unable to find a  host")
}

// AddPlayers adds a user to the existing slice of Users
func (game *Game) AddPlayers(players []*Player) {
	for _, p := range players {
		game.Players = append(game.Players, p)
	}

	return
}
