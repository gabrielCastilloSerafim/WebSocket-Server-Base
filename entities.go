package main

// Message types
const (
	DirectMessage    = 0
	BroadcastMessage = 1
)

type IncomingMessage struct {
	SenderId      string `json:"senderId"`
	DestinationId string `json:"destinationId"`
	Content       string `json:"content"`
	MessageType   int    `json:"messageType"`
}

type ResponseMessage struct {
	SenderId string `json:"senderId"`
	Content  string `json:"content"`
}

type Ping struct {
	KeepAlive bool `json:"keepAlive"`
}
