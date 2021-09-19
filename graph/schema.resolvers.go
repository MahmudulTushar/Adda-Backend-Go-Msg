package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/MahmudulTushar/Adda-Backend-Go-Msg/graph/generated"
	"github.com/MahmudulTushar/Adda-Backend-Go-Msg/graph/model"
	"github.com/MahmudulTushar/Adda-Backend-Go-Msg/repository"
)

func (r *mutationResolver) CreateNewMessage(ctx context.Context, input model.NewMessage) (*model.Message, error) {
	message := &model.NewMessage{
		Message:     input.Message,
		TimeStamp:   input.TimeStamp,
		SenderEmail: input.SenderEmail,
		RoomID:      input.RoomID,
		Name:        input.Name,
	}
	newMessage := messageRepository.Save(message)
	return &newMessage, nil
}

func (r *queryResolver) Messages(ctx context.Context) ([]*model.Message, error) {
	return messageRepository.FindAll(), nil
}

func (r *queryResolver) MessagesByRoomID(ctx context.Context, id string) ([]*model.Message, error) {
	return messageRepository.FindByRoomId(id), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var messageRepository repository.MessageRepository = repository.NewMessageRepoInstance()
