package main

//import "errors"

type GameInteractor struct {
	PlayerInteractor PlayerInteractor
	ItemRepo         ItemRepository
	GameRepo         GameRepository
	Game             Game
	Logger           Logging
	Users            []*User
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
func (gi *GameInteractor) CreateGame(host string, user *User) (Game, error) {

	player, err := gi.PlayerInteractor.CreateHost(host)
	if err != nil {
		return Game{}, err
	}

	user.Player = player

	gi.Game = NewGame("", "", user)

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

func (gi *GameInteractor) JoinGame(User *User) {
	gi.Users = append(gi.Users, User)
}
