package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// list of players for testing
var playerSlice []Player

// Fake player repo for testing
type TestPlayerRepo struct {
	player Player
}

func (pr *TestPlayerRepo) Save(player ...Player) ([]Player, []error) {
	players := make([]Player, 0)
	players = append(players, Player{Id: 1, Name: "joey", IsHost: true})

	errorList := make([]error, 0)

	return players, errorList
}
func (pr *TestPlayerRepo) FindById(id int) (Player, error) {
	if id == 1 {
		return Player{
			Id:     1,
			Name:   "joey",
			IsHost: true,
			Board:  Board{},
		}, nil
	}

	return Player{}, nil
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
	Convey("With a new player interactor", t, func() {
		playerInter := NewPlayerInteractor(&TestPlayerRepo{})

		Convey("Create a new player", func() {
			player, err := playerInter.CreateNewPlayer("joey", false)

			Convey("With a username of joey", func() {
				So(player.Name, ShouldEqual, "joey")
			})

			Convey("And is not a host", func() {
				So(player.IsHost, ShouldEqual, false)
			})

			Convey("And there is no error", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

// TestCreateNewHost creates a new player that is designated as the host
func TestCreateNewHost(t *testing.T) {
	Convey("With a new player interactor", t, func() {
		playerInter := NewPlayerInteractor(&TestPlayerRepo{})

		Convey("Create a new host", func() {
			username := "joey"
			player, err := playerInter.CreateHost(username)

			Convey("With a username of joey", func() {
				So(player.Name, ShouldEqual, username)
			})

			Convey("And is a host", func() {
				So(player.IsHost, ShouldEqual, true)
			})

			Convey("And there is no error", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

// TestCreateNewInvalidPlayer attempts to create a new player but should be an error
func TestCreateNewInvalidPlayer(t *testing.T) {
	Convey("With a new player interactor", t, func() {
		playerInter := NewPlayerInteractor(&TestPlayerRepo{})

		Convey("Create a new player with no username", func() {
			username := ""
			player, err := playerInter.CreateNewPlayer(username, false)

			Convey("Unitialized Player should be returned", func() {
				So(player, ShouldResemble, (Player{}))
			})

			Convey("Error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

// TestSavePlayers attempts to save several new players
func TestSavePlayers(t *testing.T) {
	Convey("With a new player interactor and three players", t, func() {
		playerInter := NewPlayerInteractor(&TestPlayerRepo{})
		a := Player{
			Name:   "player a",
			IsHost: false,
		}
		b := Player{
			Name:   "player b",
			IsHost: true,
		}
		c := Player{
			Name:   "player c",
			IsHost: false,
		}
		playerSlice = []Player{a, b, c}

		Convey("Save players", func() {
			players, errList := playerInter.playerRepo.Save(playerSlice...)
			So(len(errList), ShouldEqual, 0)
			So(len(players), ShouldEqual, 1)
		})
	})
}
