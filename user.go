package main

type User struct {
	Player     Player
	Connection Connection
	GameId     int
}

type UserRepository interface {
	Save(user User)
	FindById(id int) User
}

func NewUser(Connection Connection, GameId int, Player Player) *User {
	return &User{
		Connection: Connection,
		GameId:     GameId,
		Player:     Player,
	}
}
