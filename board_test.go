package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
)

// TestNewBoard creates a new board and confirms that a FREE space is already existing
func TestNewBoard(t *testing.T) {
	Convey("Given a new board", t, func() {
		board := NewBoard()

		Convey("There should be one item with a value of FREE", func() {
			So(len(board.Items), ShouldEqual, 1)
			So(board.Items[0].Value, ShouldEqual, "FREE")
		})
	})
}

// TestAddItemToBoard_NoError adds a valid item to the board
func TestAddItemToBoard_NoError(t *testing.T) {
	Convey("Given a blank board and a valid item", t, func() {
		board := NewBoard()
		item := Item{
			Id:    1,
			Value: "Says boom",
		}

		Convey("Add the item to the board", func() {
			board.Add(item)

			Convey("Board should have two items: free and the new one", func() {
				So(len(board.Items), ShouldEqual, 2)
			})
		})
	})
}

// TestAddItemToBoard_Overfill attempts to add an item to an overfilled board
func TestAddItemToBoard_Overfill(t *testing.T) {
	Convey("Given a new board", t, func() {
		board := NewBoard()

		Convey("Fill it with various items", func() {
			for i := 0; i < BOARD_SIZE-1; i++ {
				item := Item{
					Id:    i,
					Value: strconv.Itoa(i),
				}
				board.Items = append(board.Items, item)
			}

			Convey("Confirm that the board is full", func() {
				So(len(board.Items), ShouldEqual, BOARD_SIZE)
			})

			Convey("Add a new item to full board should fail", func() {
				item := Item{Id: 100, Value: "too far"}
				err := board.Add(item)
				So(err, ShouldNotBeNil)
			})
		})
	})
}

// TestAddItemToBoard_Duplicate attempts to add a duplicate item to the board
func TestAddItemToBoard_Duplicate(t *testing.T) {
	Convey("Given a new board", t, func() {
		board := NewBoard()

		Convey("Create a new item and add it to the board", func() {
			item := Item{Id: 22, Value: "duplicate"}
			err := board.Add(item)

			Convey("There should be no error and two items on the board", func() {
				So(err, ShouldBeNil)
				So(len(board.Items), ShouldEqual, 2)
			})

			Convey("Add the last item again", func() {
				err := board.Add(item)

				Convey("There should be an error and no change in item count", func() {
					So(err, ShouldNotBeNil)
					So(len(board.Items), ShouldEqual, 2)
				})
			})
		})
	})
}
