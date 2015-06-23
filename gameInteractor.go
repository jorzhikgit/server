package main

//import "errors"

type GameInteractor struct {
	PlayerInteractor PlayerInteractor
	UserRepo         UserRepository
	ItemRepo         ItemRepository
	GameRepo         GameRepository
	Game             Game
	Logger           Logging
}

func NewGameInteractor(
	pi PlayerInteractor,
	ur UserRepository,
	ir ItemRepository,
	gr GameRepository,
	log Logging) GameInteractor {

	return GameInteractor{
		PlayerInteractor: pi,
		UserRepo:         ur,
		ItemRepo:         ir,
		GameRepo:         gr,
		Logger:           log,
	}
}

// CreateGame creates a new game from the host's username
func (gi *GameInteractor) CreateGame(host string) (Game, error) {

	player := gi.PlayerInteractor.CreateHost(host)
	gi.Game = NewGame("", "", player)

	gameId, err := gi.GameRepo.Save(gi.Game)
	if err != nil {
		return Game{}, err
	}

	gi.Game.Id = gameId

	return gi.Game, nil
}
