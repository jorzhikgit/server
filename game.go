package main

import (
	"errors"
)

type Game struct {
	Id             int
	Name           string
	AvailableItems []Item // items available when creating a board
	Users          []*User
	Theme          string // not a choice, just a suggestion from the Host
}

type GameList struct {
	allGames map[int]Game
}

func NewGameList() GameList {
	return GameList{
		allGames: make(map[int]Game),
	}
}

func (rg *GameList) AddGame(Game Game) error {
	if _, ok := rg.allGames[Game.Id]; ok {
		return errors.New("Game already added")
	}

	rg.allGames[Game.Id] = Game

	return nil

}

type GameRepository interface {
	Save(game Game) (int, error)
	FindById(id int) (Game, error)

	// Player handling
	FindPlayersByGame(gameId int) ([]Player, error)
}

// Creates a new game
func NewGame(name string, theme string, host *User) Game {
	players := make([]*User, 0, 8)
	players = append(players, host)

	return Game{
		Id:             2,
		Name:           name,
		Theme:          theme,
		Users:          players,
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

func (game *Game) GetHost() (Player, error) {
	for i := 0; i < len(game.Users); i++ {
		if game.Users[i].Player.IsHost == true {
			return game.Users[i].Player, nil
		}
	}

	return Player{}, errors.New("Unable to find a  host")
}

func (game *Game) Players() (players []Player) {
	for _, user := range game.Users {
		players = append(players, user.Player)
	}
	return
}

func (game *Game) AddPlayers(users []*User) {
	for _, user := range users {
		game.Users = append(game.Users, user)
	}

	return
}
