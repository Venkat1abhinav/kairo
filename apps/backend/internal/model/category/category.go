package category

import (
	"github.com/Venkat1abhinav/kairo/internal/model"
)

type Category struct {
	model.Base
	UserID      string  `json:"userId" db:"user_id"`
	Name        string  `json:"name" db:"name"`
	Color       string  `json:"color" db:"color"`
	description *string `json:"decription" db:"description"`
}
