package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/bokuo-okubo/gqlgen-todos/entity"
	"github.com/bokuo-okubo/gqlgen-todos/graph/generated"
	"github.com/bokuo-okubo/gqlgen-todos/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	uid, err := strconv.ParseUint(input.UserID, 10, 32)
	if err != nil {
		return nil, err
	}
	record := entity.Todo{
		Text:   input.Text,
		Done:   false,
		UserID: uint(uid),
	}
	if err := r.DB.Create(&record).Error; err != nil {
		return nil, err
	}
	res := model.NewTodoFromEntity(&record)
	return res, nil
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
	var todos []entity.Todo
	r.DB.Model(&todos).Association("User")

	var rtnVal []*model.Todo
	var user model.User
	for _, t := range todos {

		rtnVal = append(rtnVal, model.NewTodoFromEntity(&t))
	}
	return rtnVal, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []entity.User
	r.DB.Find(&users)

	var rtnVal []*model.User
	for _, u := range users {
		rtnVal = append(rtnVal, model.NewUserFromEntity(&u))
	}
	return rtnVal, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
