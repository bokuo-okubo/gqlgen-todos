package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/bokuo-okubo/gqlgen-todos/entity"
	"github.com/bokuo-okubo/gqlgen-todos/graph/generated"
	"github.com/bokuo-okubo/gqlgen-todos/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	randVal, _ := rand.Int(rand.Reader, big.NewInt(100))
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", randVal),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	record := entity.User{
		Name: input.Name,
	}
	if err := r.DB.Create(&record).Error; err != nil {
		return nil, err
	}

	res := model.NewUserFromEntity(&record)

	return res, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
