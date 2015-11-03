package main

import (
	//"log"
	"net/http"
)

var AllGames GameList

func main() {

	AllGames = NewGameList()

	// create muxer
	// add route methods
	http.HandleFunc("/", index)
	http.HandleFunc("/game/create", createGame)
	http.HandleFunc("/game/join", joinGame)

	http.ListenAndServe(":8080", nil)

}

// index simply returns 200 OK
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Boyle Bingo", "Because why not")
	w.WriteHeader(200)
}

type CreateGameRequest struct {
	UserName string
	Theme    string
}

func createGame(w http.ResponseWriter, r *http.Request) {
}

func joinGame(w http.ResponseWriter, r *http.Request) {}
