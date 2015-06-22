package main

type User struct {
	Id     int
	Player Player
}

type UserRepository interface {
	Save(user User)
	FindById(id int) User
}
