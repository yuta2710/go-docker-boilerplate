package models

import (
	"time"

	"github.com/yuta_2710/go-clean-arc-reviews/modules/todo/entities"
)

type (
	InsertTodoSample struct {
		AuthId      int                   `json:"authId"`
		Title       string                `json:"title"`
		Description string                `json:"description"`
		IsCompleted bool                  `json:"isCompleted"`
		DueDate     time.Time             `json:"dueDate"`
		Priority    entities.Priority     `json:"priority"`
		Members     []InsertTodoMemberDto `json:"members"`
	}

	InsertTodoMemberDto struct {
		UserId string              `json:"userId"`
		Role   entities.MemberRole `json:"role"`
	}

	UpdateTodoSample struct {
	}

	UpdateTodoAvatarSample struct {
	}
)
