package repository

import (
	"context"
	"github.com/MahmudulTushar/Adda-Backend-Go-Msg/connection"
	"github.com/MahmudulTushar/Adda-Backend-Go-Msg/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const DatasBase = "whatsappdb"
const MessageCollection = "messagecontents"

type MessageRepository interface {
	Save(message *model.NewMessage) model.Message
	FindAll() []*model.Message
	FindByRoomId(id string) []*model.Message
	UpdateById(id string, message *model.Message) string
}

type messageRepo struct {
	client *mongo.Client
}

func (ms *messageRepo) Save(message *model.NewMessage) model.Message {
	collection := ms.client.Database(DatasBase).Collection(MessageCollection)
	_, err := collection.InsertOne(context.TODO(), message)
	if err != nil {
		log.Fatal(err)
		return model.Message{}
	}
	return model.Message{
		Message:     message.Message,
		TimeStamp:   message.TimeStamp,
		SenderEmail: message.SenderEmail,
		RoomID:      message.RoomID,
		Name:        message.Name,
	}
}

func (ms *messageRepo) FindAll() []*model.Message {
	collection := ms.client.Database(DatasBase).Collection(MessageCollection)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	var result []*model.Message
	for cursor.Next(context.TODO()) {
		var v *model.Message
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, v)
	}
	return result
}

func (ms *messageRepo) FindByRoomId(id string) []*model.Message {
	filter := bson.M{"roomId": id}
	collection := ms.client.Database(DatasBase).Collection(MessageCollection)
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	var result []*model.Message
	for cursor.Next(context.TODO()) {
		var v *model.Message
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, v)
	}
	return result
}

func (ms *messageRepo) UpdateById(id string, message *model.Message) string {
	//filter := bson.M{"_id": bson.M{"$eq": id}}
	//update := bson.M{"$set": bson.M{"received": message.Received}}
	//collection := ms.client.Database(DatasBase).Collection(MessageCollection)
	//res, err := collection.UpdateOne(
	//	context.Background(),
	//	filter,
	//	update,
	//)
	//if err != nil {
	//	log.Fatal(err)
	//	return "Error while updating the record"
	//}
	//if res.ModifiedCount > 0 {
	//	return "Record Update"
	//}
	return "No record found"
}

func NewMessageRepoInstance() MessageRepository {
	return &messageRepo{
		client: connection.MongoDBInstance.Client,
	}
}
