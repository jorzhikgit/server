package main

//import "errors"

type GameInteractor struct {
	PlayerRepo PlayerRepository
	UserRepo   UserRepository
	ItemRepo   ItemRepository
	GameRepo   GameRepository
	Logger     Logging
}

func NewGameInteractor(
	pr PlayerRepository,
	ur UserRepository,
	ir ItemRepository,
	gr GameRepository,
	log Logging) GameInteractor {

	return GameInteractor{
		PlayerRepo: pr,
		UserRepo:   ur,
		ItemRepo:   ir,
		GameRepo:   gr,
		Logger:     log,
	}
}

func (gi *GameInteractor) CreateGame(host User) error {

	return nil
}
