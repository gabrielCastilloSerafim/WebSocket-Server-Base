package main

func (server *Server) handleSendDirectMessage(incomingMessage *IncomingMessage) {
	destinationConnection := server.connections[incomingMessage.DestinationId]
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
	for _, connection := range server.connections {
		go server.sendJSON(connection, incomingMessage.DestinationId, responseMessage)
	}
}
