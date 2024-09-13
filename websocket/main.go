package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)


type WebSocketHandler struct {
	upgrader websocket.Upgrader
}

func (wsh WebSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := wsh.upgrader.Upgrade(w, r, nil);

	if err != nil {
		log.Printf("error %s when upgrading connection to websocket", err)
		return
	}

	defer c.Close()
}

func main() {
	websocketHandler := WebSocketHandler{
		upgrader: websocket.Upgrader{},
	}

	http.Handle("/", websocketHandler)
	log.Print("Starting Server...")

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}