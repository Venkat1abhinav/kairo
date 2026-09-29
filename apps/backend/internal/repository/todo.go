package repository

import (
	"context"
	"fmt"
	"uuid"

	"github.com/Venkat1abhinav/kairo/internal/model/todo"
	"github.com/Venkat1abhinav/kairo/internal/server"
	"github.com/jackc/pgx/v5"
)

type TodoRepository struct {
	server *server.Server
}

func NewTodoRepository(server *server.Server) *TodoRepository {
	return &TodoRepository{server: server}
}

func (t *TodoRepository) CreateTodo(
	ctx context.Context,
	userID string,
	payload *todo.CreateTodoPayload,
) (*todo.Todo, error) {
	stmt := `
		INSERT INTO
			todos (
				user_id,
				title,
				description,
				status,
				priority,
				due_date,
				parent_todo_id,
				category_id,
				metadata
			)
			VALUES
			(
				@user_id,
				@title,
				@description,
				@status,
				@priority,
				@due_date,
				@parent_todo_id,
				@category_id,
				@metadata
			)
	RETURNING
		*
	`

	priority := todo.PriorityMedium

	if payload.Priority != nil {
		payload.Priority = &priority
	}

	rows, err := t.server.DB.Pool.Query(ctx, stmt, pgx.NamedArgs{
		"user_id":        userID,
		"title":          payload.Title,
		"description":    payload.Description,
		"priority":       payload.Priority,
		"due_date":       payload.DueDate,
		"parent_todo_id": payload.ParentTodoID,
		"category_id":    payload.CategoryID,
		"metadata":       payload.Metadata,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to execute the todo query user_id=%s title=%s with err:%w", userID, payload.Title, err)
	}

	todoItem, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[todo.Todo])
	if err != nil {
		return nil, fmt.Errorf("failed to execute the todo query user_id=%s title=%s with err:%w", userID, payload.Title, err)
	}

	return &todoItem, nil
}

func (t *TodoRepository) GetTodoByID(ctx context.Context, userID string, todoID uuid.UUID) (*todo.PopulatedTodo, error) {
	stmt := `
		SELECT
			t.*,

			CASE
				WHEN c.id IS NOT NULL THEN to_jsonb(camel(c))
				ELSE NULL
			END AS category,

			COALESCE(
				(
					SELECT jsonb_agg(
						to_jsonb(camel(child))
						ORDER BY child.sort_order ASC, child.created_at ASC
					)
					FROM todos child
					WHERE child.parent_todo_id = t.id
					AND child.user_id = @user_id
				),
				'[]'::jsonb
			) AS children,

			COALESCE(
				(
					SELECT jsonb_agg(
						to_jsonb(camel(com))
						ORDER BY com.created_at ASC
					)
					FROM todo_comments com
					WHERE com.todo_id = t.id
					AND com.user_id = @user_id
				),
				'[]'::jsonb
			) AS comments

		FROM todos t

		LEFT JOIN todo_categories c
			ON c.id = t.category_id
			AND c.user_id = @user_id

		WHERE t.id = @id
		AND t.user_id = @user_id;
		`

	rows, err := t.server.DB.Pool.Query(ctx, stmt, pgx.NamedArgs{
		"id":      todoID,
		"user_id": userID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get the todo query user_id=%s id=%s with err:%w", userID, todoID, err)
	}

	todoItem, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[todo.PopulatedTodo])
	if err != nil {
		return nil, fmt.Errorf("failed to get the todo query user_id=%s id=%s with err:%w", userID, todoID, err)
	}

	return &todoItem, nil
}

func (t *TodoRepository) CheckTodoExists(
	ctx context.Context,
	userID string,
	todoID uuid.UUID,
) (*todo.Todo, error) {
	stmt := `
	SELECT
		*
	FROM
		TODOS
	WHERE
		id=@id
		user_id=@user_id
	`

	rows, err := t.server.DB.Pool.Query(ctx, stmt, pgx.NamedArgs{
		"user_id": userID,
		"id":      todoID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get the todo query user_id=%s id=%s with err:%w", userID, todoID, err)
	}

	todo, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[todo.Todo])
	if err != nil {
		return nil, fmt.Errorf("failed to get the todo query user_id=%s id=%s with err:%w", userID, todoID, err)
	}

	return &todo, nil
}
