// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Message struct {
	ID          string `json:"id" bson:"_id"`
	Message     string `json:"message" bson:"message"`
	Name        string `json:"name" bson:"name"`
	TimeStamp   string `json:"timeStamp" bson:"timeStamp"`
	SenderEmail string `json:"senderEmail" bson:"senderEmail"`
	RoomID      string `json:"roomId" bson:"roomId"`
}

type NewMessage struct {
	Message     string `json:"message" bson:"message"`
	Name        string `json:"name" bson:"name"`
	TimeStamp   string `json:"timeStamp" bson:"timeStamp"`
	SenderEmail string `json:"senderEmail" bson:"senderEmail"`
	RoomID      string `json:"roomId" bson:"roomId"`
}

type UpdateMessage struct {
	Received bool `json:"received"`
}