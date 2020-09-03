package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kk-no/gql-todo/graph/generated"
	"github.com/kk-no/gql-todo/graph/model"
	"github.com/kk-no/gql-todo/models"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*models.Todo, error) {
	fmt.Println("CreateTodo")
	return &models.Todo{}, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	fmt.Println("Todos")
	return []*models.Todo{}, nil
}

func (r *todoResolver) User(ctx context.Context, obj *models.Todo) (*model.User, error) {
	fmt.Println("User")
	return &model.User{}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
