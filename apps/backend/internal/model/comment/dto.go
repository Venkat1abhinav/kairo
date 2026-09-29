package comment

import (
	"uuid"

	"github.com/go-playground/validator/v10"
)

// -----------------------------------------------------------------------------

type AddCommentPayload struct {
	TodoID  uuid.UUID `param:"id" validate:"required,uuid"`
	Content string    `json:"content" validate:"required,min=1,max1000"`
}

func (p *AddCommentPayload) Validate() error {
	return validator.New().Struct(p)
}

// -----------------------------------------------------------------------------

type GetCommentsByTodoIDPayload struct {
	TodoID uuid.UUID `param:"id" validate:"requried,uuid"`
}

func (p *GetCommentsByTodoIDPayload) Validate() error {
	return validator.New().Struct(p)
}

// ------------------------------------------------------------------------------

type UpdateCommentPayload struct {
	ID      uuid.UUID `param:"id" validate:"required,uuid"`
	Content string    `json:"content" validate:"required,min=1,max1000"`
}

func (p *UpdateCommentPayload) Validate() error {
	return validator.New().Struct(p)
}

// -----------------------------------------------------------------------------

type DeleteCommentPayload struct {
	ID uuid.UUID `param:"id" validate:"required,uuid"`
}

func (p *DeleteCommentPayload) Validate() error {
	return validator.New().Struct(p)
}
