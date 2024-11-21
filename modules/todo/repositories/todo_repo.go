package repositories

import (
	"github.com/yuta_2710/go-clean-arc-reviews/modules/todo/entities"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/todo/models"
)

type TodoRepository interface {
	Insert(in *entities.Todo) (string, error)
	InsertBatch(in []*entities.Todo) error
	FindById(id string) (*entities.Todo, error)
	FindAllByUserId(userId string) ([]*entities.Todo, error)
	UpdateTodo(id string, sample *models.UpdateTodoSample) error
	UpdateAvatarOfTodo(id string, sample *models.UpdateTodoAvatarSample) error
	DeleteTodo(id string) error
}
