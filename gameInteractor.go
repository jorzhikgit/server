package main

//import "errors"

type GameInteractor struct {
	PlayerInteractor PlayerInteractor
	ItemRepo         ItemRepository
	GameRepo         GameRepository
	Game             Game
	Logger           Logging
}

func NewGameInteractor(
	pi PlayerInteractor,
	ir ItemRepository,
	gr GameRepository,
	log Logging) GameInteractor {

	return GameInteractor{
		PlayerInteractor: pi,
		ItemRepo:         ir,
		GameRepo:         gr,
		Logger:           log,
	}
}

// CreateGame creates a new game from the host's username
func (gi *GameInteractor) CreateGame(host string) (Game, error) {

	player, err := gi.PlayerInteractor.CreateHost(host)
	if err != nil {
		return Game{}, err
	}

	gi.Game = NewGame("", "", player)

	gameId, err := gi.GameRepo.Save(gi.Game)
	if err != nil {
		return Game{}, err
	}

	gi.Game.Id = gameId

	return gi.Game, nil
}

// SaveGame saves the current game by acting as a wrapper over the repository
func (gi *GameInteractor) SaveGame() error {
	_, err := gi.GameRepo.Save(gi.Game)
	if err != nil {
		return err
	}

	return nil
}
