package models

import (
	"time"

	"github.com/yuta_2710/go-clean-arc-reviews/modules/todo/entities"
)

type (
	InsertTodoSample struct {
		AuthId      int               `json:"auth_id"`
		Title       string            `json:"title"`
		Description string            `json:"description"`
		IsCompleted bool              `json:"is_completed"`
		DueDate     time.Time         `json:"due_date"`
		Priority    entities.Priority `json:"priority"`
	}

	UpdateTodoSample struct {
	}

	UpdateTodoAvatarSample struct {
	}
)
