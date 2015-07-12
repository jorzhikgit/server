package main

type User struct {
	Player Player
	Connection Connection
}

type UserRepository interface {
	Save(user User)
	FindById(id int) User
}
