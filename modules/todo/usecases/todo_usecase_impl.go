package usecases

import (
	"context"
	"fmt"

	"github.com/yuta_2710/go-clean-arc-reviews/modules/todo/entities"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/todo/models"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/todo/repositories"
	"github.com/yuta_2710/go-clean-arc-reviews/shared"
)

type TodoUsecaseImpl struct {
	Repo repositories.TodoRepository
	// AuthIdProvider shared.AuthIdProvider
}

func (tduc *TodoUsecaseImpl) Insert(ctx context.Context, in *models.InsertTodoSample) (int, error) {
	authId, ok := ctx.Value("authId").(string)

	if !ok {
		return 0, fmt.Errorf("authId not found in context")
	}

	fmt.Println("Clm no hahaha %s", authId)
	// if tduc.AuthIdProvider == nil {
	// 	return 0, fmt.Errorf("AuthIdProvider is not initialized")
	// } else {
	// 	fmt.Println("Let's go")
	// }

	decodedId, _ := shared.DecomposeUidV2(string(authId))
	fmt.Println(in.Members)

	// in.Priority = priority

	for i, _ := range in.Members {
		in.Members[i].TodoId = 1
	}

	toEntity := &entities.Todo{
		// AuthId:      authId,
		UserId:      int(decodedId.GetLocalID()),
		Title:       in.Title,
		Description: in.Description,
		IsCompleted: in.IsCompleted,
		DueDate:     in.DueDate,
		Priority:    in.Priority,
		// Members:     in.Members,
	}
	fmt.Println("Todo entity data inserted")

	newTodoId, err := tduc.Repo.InsertTodo(toEntity)

	if err != nil {
		panic("Cannot insert todo")
	}

	members := []entities.TodoMember{}

	for _, mem := range in.Members {
		// decodedId, _ := shared.DecomposeUidV2(string(authId))

		members = append(members, entities.TodoMember{
			UserId: mem.UserId,
			TodoId: newTodoId,
			Role:   mem.Role,
		})
	}

	if len(members) > 0 {
		err = tduc.Repo.InsertTodoMembers(newTodoId, members)

		if err != nil {
			return 0, fmt.Errorf("failed to insert todo members: %v", err)
		}
	}

	// Update members to database

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
		// AuthIdProvider: authIdProvider,
	}
}
