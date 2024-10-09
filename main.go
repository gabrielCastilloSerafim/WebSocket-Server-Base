package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	server := Server{
		connections: make(map[string]*websocket.Conn),
	}
	http.Handle("/", websocket.Handler(server.stablishWSConnection))
	port := "1337"
	fmt.Printf("🌎 Server listening on port: %s 🌎\n", port)
	panic(http.ListenAndServe(":"+port, nil))
}
