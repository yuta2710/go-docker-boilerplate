package usecases

import (
	"context"

	"github.com/yuta_2710/go-clean-arc-reviews/modules/todo/entities"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/todo/models"
)

type TodoUsecase interface {
	Insert(ctx context.Context, in *models.InsertTodoSample) (string, error)
	InsertBatch(ctx context.Context, in []*models.InsertTodoSample) error
	FindById(ctx context.Context, id string) (*entities.Todo, error)
	FindAllByUserId(ctx context.Context, userId string) ([]*entities.Todo, error)
	UpdateTodo(ctx context.Context, id string, sample *models.UpdateTodoSample) error
	UpdateAvatarOfTodo(ctx context.Context, id string, sample *models.UpdateTodoAvatarSample) error
	DeleteTodo(ctx context.Context, id string) error
}
