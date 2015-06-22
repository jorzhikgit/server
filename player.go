package main

type Player struct {
	Id int
	Name string
	IsHost bool
	Board Board
}

type PlayerRepository interface {
	Save(player Player)
	FindById(id int) Player
}

