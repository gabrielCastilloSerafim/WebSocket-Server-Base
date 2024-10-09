package main

import (
	"io"

	"golang.org/x/net/websocket"
)

type Server struct {
	connections map[string]*websocket.Conn
}

func (server *Server) stablishWSConnection(currentConnection *websocket.Conn) {
	header := currentConnection.Request().Header
	connectionId := header.Get("userId")
	server.connections[connectionId] = currentConnection
	server.startListening(currentConnection, connectionId)
}

func (server *Server) startListening(currentConnection *websocket.Conn, connectionId string) {
	incomingMessage := &IncomingMessage{}
	for {
		err := websocket.JSON.Receive(currentConnection, incomingMessage)
		if err != nil && err == io.EOF {
			currentConnection.Close()
			server.connections[connectionId] = nil
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
