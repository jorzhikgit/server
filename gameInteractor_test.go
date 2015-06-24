package main

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Item repository for testing
type TestItemRepo struct{}

func (ir *TestItemRepo) Save(item Item) {
}
func (ir *TestItemRepo) FindById(id int) Item {
	return Item{}
}

// Bad game repository, only returns errors
type BadGameRepo struct {
	game Game
}

func (gr *BadGameRepo) Save(game Game) (int, error) { return 0, errors.New("Unable to save game") }
func (gr *BadGameRepo) FindById(id int) (Game, error) {
	return Game{}, errors.New("Unable to find game")
}
func (gr *BadGameRepo) FindPlayersByGame(gameId int) ([]Player, error) { return make([]Player, 0), nil }

// Game repository for testing
type TestGameRepo struct {
	game Game
}

func (gr *TestGameRepo) Save(game Game) (int, error) { return 0, nil }
func (gr *TestGameRepo) FindById(id int) (Game, error) {
	return Game{}, nil
}
func (gr *TestGameRepo) FindHost() Player {
	return Player{
		Id:     1,
		Name:   "joey",
		IsHost: true,
		Board:  Board{},
	}
}
func (gr *TestGameRepo) FindPlayersByGame(gameId int) ([]Player, error) { return make([]Player, 0), nil }

// Logging for testing
type TestLogger struct{}

func (log *TestLogger) Log(message string) error {
	return errors.New(message)
}

func setupGameInteractor() GameInteractor {
	playerInteractor := NewPlayerInteractor(&TestPlayerRepo{})
	gameRepo := TestGameRepo{}
	itemRepo := TestItemRepo{}
	logger := TestLogger{}

	return NewGameInteractor(playerInteractor, &itemRepo, &gameRepo, &logger)
}

// TestNewGameInteractor generates a new GameInteractor from inputs
func TestNewGameInteractor(t *testing.T) {

	Convey("Create a new GameInteractor", t, func() {
		playerInteractor := NewPlayerInteractor(&TestPlayerRepo{})
		gameRepo := TestGameRepo{}
		itemRepo := TestItemRepo{}
		logger := TestLogger{}

		gameInter := NewGameInteractor(playerInteractor, &itemRepo, &gameRepo, &logger)

		Convey("Should be equal to custom GameInteractor", func() {
			customGame := GameInteractor{
				PlayerInteractor: playerInteractor,
				GameRepo:         &gameRepo,
				ItemRepo:         &itemRepo,
				Logger:           &logger,
			}
			So(gameInter, ShouldResemble, customGame)
		})
	})
}

// TestCreateGame creates a new game using a game interactor and host
func TestCreateGame(t *testing.T) {
	Convey("Using a new GameInteractor and a new host", t, func() {
		gameInter := setupGameInteractor()
		user := User{
			Id: 1,
			Player: Player{
				Id:     1,
				Name:   "joey",
				IsHost: true,
			},
		}

		Convey("Create a new game", func() {
			_, err := gameInter.CreateGame(user.Player.Name)

			Convey("And have the host be joey", func() {
				So(err, ShouldBeNil)
				So(gameInter.Game.Players[0], ShouldResemble, user.Player)
			})
		})

	})
}

// TestGetHost get's the Player that is the host
func TestGetHost(t *testing.T) {
	Convey("Using a new GameInteractor and a new host", t, func() {
		gameInter := setupGameInteractor()
		username := "joey"

		Convey("Create a new game", func() {
			game, err := gameInter.CreateGame(username)

			Convey("Where the game has not name", func() {
				So(err, ShouldBeNil)
				So(game.Name, ShouldEqual, "")
			})

			Convey("And the host is username", func() {
				host, err := game.GetHost()
				So(err, ShouldBeNil)
				So(host.Name, ShouldEqual, username)
			})
		})
	})
}

// TestSaveGame saves the game to the database
func TestSaveGame(t *testing.T) {
	Convey("Using a new GameInteractor and a host username", t, func() {
		gameInter := setupGameInteractor()
		username := "joey"

		Convey("Create a new game", func() {
			game, err := gameInter.CreateGame(username)

			Convey("Where the game has no name", func() {
				So(err, ShouldBeNil)
				So(game.Name, ShouldEqual, "")
			})

			Convey("Save the game", func() {
				err := gameInter.SaveGame()
				So(err, ShouldBeNil)
			})
		})
	})
}
