package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// TestNewGame creates a new game
func TestNewGame(t *testing.T) {
	Convey("Given a new game", t, func() {
		host := Player{
			Id:     1,
			Name:   "joey",
			IsHost: true,
			Board:  Board{},
		}
		game := NewGame("Barry Bingo", "All the funny things Barry does", host)

		Convey("There should be a host", func() {
			host := Player{}
			for _, player := range game.Players {
				if player.IsHost == true {
					host = player
				}
			}
			So(host, ShouldNotResemble, Player{})
		})

		Convey("There should be a name", func() {
			So(game.Name, ShouldNotBeBlank)
		})

		Convey("There should be a theme", func() {
			So(game.Theme, ShouldNotBeBlank)
		})

		Convey("There should be no items", func() {
			So(len(game.AvailableItems), ShouldEqual, 0)
		})
	})
}

// TestAddToAvailableItems adds a new item to the list of available itmes
func TestAddToAvailableItems(t *testing.T) {
	Convey("Given a new game", t, func() {
		host := Player{
			Id: 1,
			Name: "joey",
			IsHost: true,
			Board: Board{},
		}
		game := NewGame("Barry Bingo", "All the funny things Barry does", host)

		Convey("Add a new item to the available items", func() {
			item := Item {Id: 1, Value: "fudge"}
			count, err := game.AddToAvailable(item)

			Convey("There should be one item and no errors", func() {
				So(count, ShouldEqual, 1)
				So(err, ShouldBeNil)
			})
		})
	})
}

// TestAddToAvailableItems_Duplicate adds the same item twice
func TestAddToAvailableItems_Duplicate(t *testing.T) {
	Convey("Given a new game", t, func() {
		host := Player {
			Id: 1,
			Name: "joey",
			IsHost: true,
			Board: Board{},
		}
		game := NewGame("Barry Bingo", "stuff", host)

		Convey("Add a new item to the available items", func() {
			item := Item {Id: 1, Value: "fudge"}
			count, err := game.AddToAvailable(item)

			Convey("There should be one item and no errors", func() {
				So(count, ShouldEqual, 1)
				So(err, ShouldBeNil)
			})

			Convey("Create a new item with the same value and add to list", func() {
				item2 := Item{Id: 2, Value: "fudge"}
				count, err = game.AddToAvailable(item2)

				Convey("There should be one item and an error", func() {
					So(count, ShouldEqual, 1)
					So(err, ShouldNotBeNil)
				})
			})
		})
	})
}
