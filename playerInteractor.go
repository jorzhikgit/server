package main

import (
	"errors"
)

type PlayerInteractor struct {
	playerRepo PlayerRepository
}

// NewPlayerInteractor creates a new PlayerInteractor
func NewPlayerInteractor(pr PlayerRepository) PlayerInteractor {
	return PlayerInteractor{
		playerRepo: pr,
	}
}

// CreateNewPlayer creates a new player using a username and a flag to determine
// host status.
// An empty player is returned if there is an error
func (p *PlayerInteractor) CreateNewPlayer(name string, isHost bool) Player {
	player := Player{
		Name:   name,
		IsHost: isHost,
	}

	savedPlayers, errList := p.playerRepo.Save(player)
	if len(errList) > 0 {
		return Player{}
	}

	player.Id = savedPlayers[0].Id

	return player
}

// CreateHost takes a username and creates a new host. This is a shortcut
// to `CreateNewPlayer` to reduce magical values throughout the codebase.
// An empty Player is returned if there is an error
func (p *PlayerInteractor) CreateHost(name string) Player {
	return p.CreateNewPlayer(name, true)
}

// SavePlayers takes a slice of players and saves them all through the
// repository.
func (p *PlayerInteractor) SavePlayers(Players ...Player) error {
	_, err := p.playerRepo.Save(Players...)
	if len(err) > 0 {
		return errors.New("Unable to save some players")
	}

	return nil
}
