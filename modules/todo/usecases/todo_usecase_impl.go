package usecases

import (
	"context"
	"fmt"

	"github.com/yuta_2710/go-clean-arc-reviews/modules/todo/entities"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/todo/models"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/todo/repositories"
)

type TodoUsecaseImpl struct {
	Repo repositories.TodoRepository
}

func (tduc *TodoUsecaseImpl) Insert(ctx context.Context, in *models.InsertTodoSample) (string, error) {
	authId, ok := ctx.Value("authId").(string)
	fmt.Println(authId, ok)
	if !ok {
		return "", fmt.Errorf("authId not found in context")
	}

	toEntity := &entities.Todo{
		AuthId:      authId,
		Title:       in.Title,
		Description: in.Description,
		IsCompleted: in.IsCompleted,
		DueDate:     in.DueDate,
		Priority:    in.Priority,
	}

	fmt.Println("Todo entity data inserted")
	fmt.Println(toEntity)

	newTodoId, err := tduc.Repo.Insert(toEntity)

	if err != nil {
		panic("Cannot insert todo")
	}

	return newTodoId, nil
}

func (tduc *TodoUsecaseImpl) InsertBatch(ctx context.Context, in []*models.InsertTodoSample) error {
	return nil
}

func (tduc *TodoUsecaseImpl) FindById(ctx context.Context, id string) (*entities.Todo, error) {
	return nil, nil
}

func (tduc *TodoUsecaseImpl) FindAllByUserId(ctx context.Context, userId string) ([]*entities.Todo, error) {
	todos, err := tduc.Repo.FindAllByUserId(userId)

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (tduc *TodoUsecaseImpl) UpdateTodo(ctx context.Context, id string, sample *models.UpdateTodoSample) error {
	return nil
}

func (tduc *TodoUsecaseImpl) UpdateAvatarOfTodo(ctx context.Context, id string, sample *models.UpdateTodoAvatarSample) error {
	return nil
}

func (tduc *TodoUsecaseImpl) DeleteTodo(ctx context.Context, id string) error {
	return nil
}

func NewTodoUsecaseImpl(repo repositories.TodoRepository) TodoUsecase {
	return &TodoUsecaseImpl{
		Repo: repo,
	}
}
