package main

func (server *Server) handleSendDirectMessage(incomingMessage *IncomingMessage) {
	server.Lock()
	destinationConnection := server.connections[incomingMessage.DestinationId]
	server.Unlock()
	if destinationConnection == nil {
		return
	}
	responseMessage := ResponseMessage{
		SenderId: incomingMessage.SenderId,
		Content:  incomingMessage.Content,
	}
	server.sendJSON(destinationConnection, incomingMessage.DestinationId, responseMessage)
}

func (server *Server) handleSendBrodcastMessage(incomingMessage *IncomingMessage) {
	responseMessage := ResponseMessage{
		SenderId: incomingMessage.SenderId,
		Content:  incomingMessage.Content,
	}
	server.Lock()
	serverConnections := server.connections
	server.Unlock()
	for _, connection := range serverConnections {
		go server.sendJSON(connection, incomingMessage.DestinationId, responseMessage)
	}
}
