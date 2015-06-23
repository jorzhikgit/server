package main

type PlayerInteractor struct {
	playerRepo PlayerRepository
}

func NewPlayerInteractor(pr PlayerRepository) PlayerInteractor {
	return PlayerInteractor{
		playerRepo: pr,
	}
}

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

func (p *PlayerInteractor) CreateHost(name string) Player {
	return p.CreateNewPlayer(name, true)
}
