package comment

import (
	"github.com/Venkat1abhinav/kairo/internal/model"
	"github.com/google/uuid"
)

type Comment struct {
	model.Base
	TodoID  uuid.UUID `json:"todoId" db:"todo_id"`
	UserID  string    `json:"userId" db:"user_id"`
	Content string    `json:"content" db:"content"`
}
