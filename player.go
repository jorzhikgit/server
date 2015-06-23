package main

type Player struct {
	Id     int
	Name   string
	IsHost bool
	Board  Board
}

type PlayerRepository interface {
	Save(player Player) (int, error) // return an ID number for the saved player
	FindById(id int) (Player, error)
}
