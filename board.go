package main

import (
	"errors"
)

const BOARD_SIZE int = 25

type Board struct {
	Id    int
	Items []Item // items on a board
}

type BoardRepository interface {
	Save(board Board)
	FindById(id int) Board
}

// Creates a new board
func NewBoard() Board {
	freeItem := Item{
		Id:    0,
		Value: "FREE",
	}

	items := make([]Item, 0, 25)
	items = append(items, freeItem)

	return Board{
		Id:    2,
		Items: items,
	}
}

// Add manages teh rules for adding an item to a board
func (board *Board) Add(item Item) error {
	// do not overfill the board
	if len(board.Items) >= BOARD_SIZE {
		return errors.New("Game board is at maximum capacity")
	}

	// do not allow duplicate items
	for _, existingItem := range board.Items {
		if existingItem.Id == item.Id {
			return errors.New("Game board can not have duplicate items")
		}
	}

	board.Items = append(board.Items, item)
	return nil
}
