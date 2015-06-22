package main

type Item struct {
	Id    int
	Value string
}

type ItemRepository interface {
	Save(item Item)
	FindById(id int) Item
}
