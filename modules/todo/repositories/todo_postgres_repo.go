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
	fmt.Println("\nCai lon")
	fmt.Println(in)
	result := tdr.db.GetDb().Create(in)

	if result.Error != nil {
		fmt.Println("Error inserting")
		panic(result.Error.Error())
	}

	// TODO: Mask the ID of the todo
	fmt.Println("[INSERTED DATA SUCCESSFULLY]")
	in.Mask(shared.DbTypeTodo)

	return "", nil
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

func NewTodoPostgresRepository(db database.Database) TodoRepository {
	return &TodoPostgresRepository{
		db: db,
	}
}
