package main

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Player repository for testing
type TestPlayerRepo struct{}

func (pr *TestPlayerRepo) Save(player Player) {
}
func (pr *TestPlayerRepo) FindById(id int) Player {
	return Player{}
}

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
type TestGameRepo struct{}

func (gr *TestGameRepo) Save(game Game) {
}
func (gr *TestGameRepo) FindById(id int) Game {
	return Game{}
}

// Logging for testing
type TestLogger struct{}

func (log *TestLogger) Log(message string) error{
	return errors.New(message)
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
			customGame := GameInteractor {
				PlayerRepo: &playerRepo,
				GameRepo: &gameRepo,
				UserRepo: &userRepo,
				ItemRepo: &itemRepo,
				Logger: &logger,
			}
			So(gameInter, ShouldResemble, customGame)
		})
	})
}
