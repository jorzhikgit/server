package main

import (
	"log"
	"net/http"

	evbus "github.com/asaskevich/EventBus"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// Global game related objects
var currentGames GameList
var gameHub Hub
var events *evbus.EventBus

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	//conn, err := upgrader.Upgrade(w, r, nil)
	//if err != nil {
	//log.Println(err)
	//return
	//}

	// create a connection struct
	//c := &connection{send: make(chan []byte, 256), ws: conn}
	// register with hub
	// go writer
	// run reader
}

func main() {
	// create db connection

	// event bus
	events = evbus.New()

	// run hub
	gameHub = NewHub()
	go gameHub.Run()

	// currently running games
	currentGames = NewGameList()

	http.HandleFunc("/", rootHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
