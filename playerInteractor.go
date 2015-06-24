package main

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

	newId, err := p.playerRepo.Save(player)
	if err != nil {
		return Player{}
	}

	player.Id = newId

	return player
}

// CreateHost takes a username and creates a new host. This is a shortcut
// to `CreateNewPlayer` to reduce magical values throughout the codebase.
// An empty Player is returned if there is an error
func (p *PlayerInteractor) CreateHost(name string) Player {
	return p.CreateNewPlayer(name, true)
}
