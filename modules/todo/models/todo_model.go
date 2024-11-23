package models

import (
	"time"

	"github.com/yuta_2710/go-clean-arc-reviews/modules/todo/entities"
)

type (
	InsertTodoSample struct {
		AuthId      string                `json:"authId"`
		Title       string                `json:"title"`
		Description string                `json:"description"`
		IsCompleted bool                  `json:"isCompleted"`
		DueDate     time.Time             `json:"dueDate"`
		Priority    entities.Priority     `json:"priority"`
		Members     []entities.TodoMember `json:"members"`
	}

	InsertTodoMemberDto struct {
		UserId int                 `json:"userId"`
		TodoId int                 `json:"todoId"`
		Role   entities.MemberRole `json:"role"`
	}

	UpdateTodoSample struct {
	}

	UpdateTodoAvatarSample struct {
	}
)
