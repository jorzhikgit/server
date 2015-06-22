package main

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// User repository for testing
type TestUserRepo struct{}

func (ur *TestUserRepo) Save(user User) {
}
func (ur *TestUserRepo) FindById(id int) User {
	return User{}
}

// Item repository for testing
type TestItemRepo struct{}

func (ir *TestItemRepo) Save(item Item) {
}
func (ir *TestItemRepo) FindById(id int) Item {
	return Item{}
}

// Game repository for testing
type TestGameRepo struct {
	game Game
}

func (gr *TestGameRepo) Save(game Game) {
}
func (gr *TestGameRepo) FindById(id int) Game {
	return Game{}
}
func (gr *TestGameRepo) FindHost() Player {
	for _, p := range gr.game.Players {
		if p.IsHost == true {
			return p
		}
	}

	return Player{}
}

// Logging for testing
type TestLogger struct{}

func (log *TestLogger) Log(message string) error {
	return errors.New(message)
}

func setupGameInteractor() GameInteractor {
	userRepo := TestUserRepo{}
	playerRepo := TestPlayerRepo{}
	gameRepo := TestGameRepo{}
	itemRepo := TestItemRepo{}
	logger := TestLogger{}

	return NewGameInteractor(&playerRepo, &userRepo, &itemRepo, &gameRepo, &logger)
}

// TestNewGameInteractor generates a new GameInteractor from inputs
func TestNewGameInteractor(t *testing.T) {

	Convey("Create a new GameInteractor", t, func() {
		userRepo := TestUserRepo{}
		playerRepo := TestPlayerRepo{}
		gameRepo := TestGameRepo{}
		itemRepo := TestItemRepo{}
		logger := TestLogger{}

		gameInter := NewGameInteractor(&playerRepo, &userRepo, &itemRepo, &gameRepo, &logger)

		Convey("Should be equal to custom GameInteractor", func() {
			customGame := GameInteractor{
				PlayerRepo: &playerRepo,
				GameRepo:   &gameRepo,
				UserRepo:   &userRepo,
				ItemRepo:   &itemRepo,
				Logger:     &logger,
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
			err := gameInter.CreateGame(user)

			Convey("And have the host be joey", func() {
				So(err, ShouldBeNil)
				So(gameInter.GameRepo.FindHost(), ShouldResemble, user.Player)
			})
		})

	})
}
