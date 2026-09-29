package category

import (
	"uuid"

	"github.com/go-playground/validator/v10"
)

// -------------------------------------------------------------------------------------------

type CreateCategoryPayload struct {
	Name        string  `json:"name" validate:"required,min=1,max=255"`
	Color       string  `json:"color" validate:"required,hexcolor"`
	Description *string `json:"description" validate:"omitempty,min=1,max=1000"`
}

func (p *CreateCategoryPayload) Validate() error {
	return validator.New().Struct(p)
}

// -------------------------------------------------------------------------------------------

type UpdateCategoryPayload struct {
	ID          uuid.UUID `param:"id" validate:"required,uuid"`
	Name        *string   `json:"name" validate:"omitempty,min=1,max=255"`
	Color       *string   `json:"color" validate:"omitempty,hexcolor"`
	Description *string   `json:"description" validate:"omitempty,min=1,max=1000"`
}

func (p *UpdateCategoryPayload) Validate() error {
	return validator.New().Struct(p)
}

// -------------------------------------------------------------------------------------------

type GetCategoriesQuery struct {
	Page   *int    `json:"page" validate:"omitempty,min=1"`
	Limit  *int    `json:"limit" validate:"omitempty,min=1,max=100"`
	Sort   *string `json:"sort" validate:"omitempty,oneof=created_at updated_at name"`
	Order  *string `json:"order" validate:"omitempty,oneof=asc desc"`
	Search *string `json:"search" validate:"omitempty,min=1"`
}

func (q *GetCategoriesQuery) Validate() error {
	if err := validator.New().Struct(q); err != nil {
		return err
	}

	if q.Page == nil {
		defaultPage := 1
		q.Page = &defaultPage
	}

	if q.Limit == nil {
		defaultLimit := 50
		q.Limit = &defaultLimit
	}

	if q.Sort == nil {
		defaultSort := "created_at"
		q.Sort = &defaultSort
	}

	if q.Order == nil {
		defaultOrder := "asc"
		q.Order = &defaultOrder
	}

	return nil
}
