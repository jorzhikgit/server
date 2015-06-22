package main

type PlayerInteractor struct {
	playerRepo PlayerRepository
}

func NewPlayerInteractor(pr PlayerRepository) PlayerInteractor {
	return PlayerInteractor{
		playerRepo: pr,
	}
}
