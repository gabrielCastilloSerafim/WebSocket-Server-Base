package main

import "golang.org/x/net/websocket"

func (server *Server) sendJSON(destinationConnection *websocket.Conn, destinationConnectionId string, responseMessage ResponseMessage) {
	err := websocket.JSON.Send(destinationConnection, responseMessage)
	if err != nil {
		server.connections[destinationConnectionId] = nil
	}
}
