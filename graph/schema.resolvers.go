package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"question-bank-service/database"
	"question-bank-service/graph/generated"
	"question-bank-service/graph/model"
)

func (r *mutationResolver) CreateQuestion(ctx context.Context, input model.NewQuestion) (*model.Question, error) {
	return db.InsertQuestionById(input), nil
}

func (r *queryResolver) Question(ctx context.Context, id string) (*model.Question, error) {
	return db.FindQuestionById(id), nil
}

func (r *queryResolver) Questions(ctx context.Context) ([]*model.Question, error) {
	return db.Questions(), nil
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
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
var db = database.Connect("mongodb+srv://Tantely:QraeYdSp35HaVGY@mongodb-playground.32w54.mongodb.net/?retryWrites=true&w=majority")
