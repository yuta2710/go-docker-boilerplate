package repositories

import (
	"fmt"

	"github.com/yuta_2710/go-clean-arc-reviews/database"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/todo/entities"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/todo/models"
	"github.com/yuta_2710/go-clean-arc-reviews/shared"
)

type TodoPostgresRepository struct {
	db database.Database
}

func (tdr *TodoPostgresRepository) Insert(in *entities.Todo) (string, error) {
	// TODO: Insert do database
	tx := tdr.db.GetDb().Begin()

	result := tx.Create(in)

	if result.Error != nil {
		fmt.Println("Error inserting")
		tx.Rollback()
		return "", fmt.Errorf("failed to insert todo: %v", result.Error)
	}

	// Create associates members
	if len(in.Members) > 0 {
		for _, member := range in.Members {
			member.TodoId = in.FakeId.String()
			if err := tx.Create(&member).Error; err != nil {
				tx.Rollback()
				return "", fmt.Errorf("failed to insert todo member: %v", err)
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	// TODO: Mask the ID of the todo
	fmt.Println("[INSERTED DATA SUCCESSFULLY]")
	in.Mask(shared.DbTypeTodo)

	return in.FakeId.String(), nil
}

func (tdr *TodoPostgresRepository) InsertBatch(in []*entities.Todo) error {
	return nil
}

func (tdr *TodoPostgresRepository) FindById(id string) (*entities.Todo, error) {
	return nil, nil
}

func (tdr *TodoPostgresRepository) FindAllByUserId(userId string) ([]*entities.Todo, error) {
	var todos []*entities.Todo
	result := tdr.db.GetDb().Where("auth_id = ?", userId).Find(&todos)

	if result.Error != nil {
		return nil, result.Error
	}
	return todos, nil
}

func (tdr *TodoPostgresRepository) UpdateTodo(id string, sample *models.UpdateTodoSample) error {
	return nil
}

func (tdr *TodoPostgresRepository) UpdateAvatarOfTodo(id string, sample *models.UpdateTodoAvatarSample) error {
	return nil
}

func (tdr *TodoPostgresRepository) DeleteTodo(id string) error {
	return nil
}

func (tdr *TodoPostgresRepository) AddUserForTodo(userId string) error {
	return nil
}

func NewTodoPostgresRepository(db database.Database) TodoRepository {
	return &TodoPostgresRepository{
		db: db,
	}
}
