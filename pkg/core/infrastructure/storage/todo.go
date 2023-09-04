package storage

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/murasame29/echo-hex-arch-template/pkg/core/entities/request"
	"github.com/murasame29/echo-hex-arch-template/pkg/core/entities/response"
)

// ここではTodoに関するいわゆるリポジトリ層を実装する

type TodoStorage interface {
	ListTodo(arg request.ListTodo) ([]response.ListTodo, error)
	CreateTodo(arg request.CreateTodo) (response.Todo, error)
	DeleteAllTodo(arg request.DeleteAllTodo) error
	GetTodo(todoID string) (response.Todo, error)
	UpdateTodo(arg request.UpdateTodo) (response.Todo, error)
	DeleteTodo(userID string) error
}

type todoStorage struct {
	db  *sql.DB
	ctx context.Context
}

func NewTodoStorage(ctx context.Context, db *sql.DB) TodoStorage {
	return &todoStorage{
		db:  db,
		ctx: ctx,
	}
}

func (ts *todoStorage) ListTodo(arg request.ListTodo) ([]response.ListTodo, error) {
	query := `SELECT todo_id,title,created_at,is_complete FROM todos WHERE user_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`

	rows, err := ts.db.QueryContext(ts.ctx, query, arg.UserID, arg.PageSize, (arg.PageID-1)*arg.PageSize)
	if err != nil {
		return nil, err
	}

	var todos []response.ListTodo

	for rows.Next() {
		var todo response.ListTodo
		err = rows.Scan(&todo.TodoID, &todo.Title, &todo.CreatedAt, &todo.IsComplete)

		// true  => complateのタスクを含めない
		// false => complateのタスクを含める
		if arg.IsComplete {
			if !todo.IsComplete {
				todos = append(todos, todo)
			}
		} else {
			todos = append(todos, todo)
		}
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (ts *todoStorage) CreateTodo(arg request.CreateTodo) (response.Todo, error) {
	query := `INSERT INTO todos(todo_id,user_id,title,description)VALUES($1,$2,$3,$4)RETURNING todo_id,title,description,created_at,updated_at,is_complete`
	rows := ts.db.QueryRowContext(ts.ctx, query, uuid.New().String(), arg.UserID, arg.Title, arg.Description)

	var todo response.Todo
	err := rows.Scan(&todo.TodoID, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt, &todo.IsComplete)
	return todo, err
}

func (ts *todoStorage) DeleteAllTodo(arg request.DeleteAllTodo) error {
	query := `DELETE FROM todos WHERE user_id = $1`
	_, err := ts.db.ExecContext(ts.ctx, query, arg.UserID)
	return err
}

func (ts *todoStorage) GetTodo(todoID string) (response.Todo, error) {
	query := `SELECT todo_id,title,description,created_at,updated_at,is_complete FROM todos WHERE todo_id = $1`
	rows := ts.db.QueryRowContext(ts.ctx, query, todoID)

	var todo response.Todo
	err := rows.Scan(&todo.TodoID, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt, &todo.IsComplete)
	return todo, err
}

func (ts *todoStorage) UpdateTodo(arg request.UpdateTodo) (response.Todo, error) {
	query := `UPDATE todos SET title=$1,description=$2,updated_at=$3,is_complete=$4 WHERE todo_id = $5 RETURNING todo_id,title,description,created_at,updated_at,is_complete`
	rows := ts.db.QueryRowContext(ts.ctx, query,
		arg.Title,
		arg.Description,
		time.Now(),
		arg.IsComplete,
		arg.TodoID,
	)

	var todo response.Todo
	err := rows.Scan(&todo.TodoID, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt, &todo.IsComplete)
	return todo, err
}

func (ts *todoStorage) DeleteTodo(userID string) error {
	query := `DELETE FROM todos WHERE todo_id = $1`
	_, err := ts.db.ExecContext(ts.ctx, query, userID)
	return err
}
