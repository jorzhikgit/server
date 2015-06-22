package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Fake player repo for testing
type TestPlayerRepo struct {
	player Player
}

func (pr *TestPlayerRepo) Save(player Player) int {
	return 1
}
func (pr *TestPlayerRepo) FindById(id int) Player {
	if id == 1 {
		return Player{
			Id:     1,
			Name:   "joey",
			IsHost: true,
			Board:  Board{},
		}
	}

	return Player{}
}

// TestNewPlayerInteractor creates a new object to manage players
func TestNewPlayerInteractor(t *testing.T) {
	Convey("Create a new player interactor", t, func() {
		playerInter := NewPlayerInteractor(&TestPlayerRepo{})
		Convey("Where the PlayerInteractor matches custom PlayerInteractor", func() {
			customInter := PlayerInteractor{playerRepo: &TestPlayerRepo{}}
			So(playerInter, ShouldResemble, customInter)
		})
	})
}

// TestCreateNewPlayer creates a new player
func TestCreateNewPlayer(t *testing.T) {
}

// TestCreateNewInvalidPlayer attempts to create a new player but should be an error
func TestCreateNewInvalidPlayer(t *testing.T) {
}
