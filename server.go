package main

import (
	"io"
	"sync"

	"golang.org/x/net/websocket"
)

type Server struct {
	sync.Mutex
	connections map[string]*websocket.Conn
}

func NewServer() *Server {
	server := Server{
		connections: make(map[string]*websocket.Conn),
	}
	return &server
}

func (server *Server) stablishWSConnection(currentConnection *websocket.Conn) {
	header := currentConnection.Request().Header
	connectionId := header.Get("userId")
	server.Lock()
	server.connections[connectionId] = currentConnection
	server.Unlock()
	server.startListening(currentConnection, connectionId)
}

func (server *Server) startListening(currentConnection *websocket.Conn, connectionId string) {
	incomingMessage := &IncomingMessage{}
	for {
		err := websocket.JSON.Receive(currentConnection, incomingMessage)
		if err != nil && err == io.EOF {
			currentConnection.Close()
			server.Lock()
			delete(server.connections, connectionId)
			server.Unlock()
			break
		}
		switch incomingMessage.MessageType {
		case DirectMessage:
			server.handleSendDirectMessage(incomingMessage)
		case BroadcastMessage:
			server.handleSendBrodcastMessage(incomingMessage)
		}
	}
}
