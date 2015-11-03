package main

import "errors"

// GameList represents all of the currently running games
type GameList struct {
	allGames map[int]Game
}

// NewGameList generates a new GameList to work with
func NewGameList() GameList {
	return GameList{
		allGames: make(map[int]Game),
	}
}

// AddGame adds a new game to the list of currently running games
func (rg *GameList) AddGame(Game Game) error {
	if _, ok := rg.allGames[Game.Id]; ok {
		return errors.New("Game already added")
	}

	rg.allGames[Game.Id] = Game

	return nil
}
