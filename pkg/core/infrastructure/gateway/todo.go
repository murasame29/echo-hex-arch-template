package gateway

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/murasame29/echo-hex-arch-template/pkg/internal/helpers/token"

	"github.com/murasame29/echo-hex-arch-template/pkg/core/entities/request"
	"github.com/murasame29/echo-hex-arch-template/pkg/core/entities/response"
	"github.com/murasame29/echo-hex-arch-template/pkg/core/infrastructure/storage"
)

type TodoGateway interface {
	ListUser(body request.ListTodo, payload *token.Payload) (int, []response.ListTodo)
	CreateTodo(body request.CreateTodo, payload *token.Payload) (int, response.Todo)
	DeleteAllTodo(body request.DeleteAllTodo, payload *token.Payload) (int, string)
	GetTodo(todoID string) (int, response.Todo)
	UpdateTodo(body request.UpdateTodo, payload *token.Payload) (int, response.Todo)
	DeleteTodo(todoID string) (int, string)
}

type TodoLogic struct {
	storage storage.TodoStorage
}

func NewTodoGateway(ctx context.Context, db *sql.DB) TodoGateway {
	return &TodoLogic{storage.NewTodoStorage(ctx, db)}
}

func (tl *TodoLogic) ListUser(body request.ListTodo, payload *token.Payload) (int, []response.ListTodo) {
	if payload.UserID != body.UserID {
		return http.StatusUnauthorized, nil
	}

	result, err := tl.storage.ListTodo(body)
	log.Println(err)
	if err != nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, result
}

func (tl *TodoLogic) CreateTodo(body request.CreateTodo, payload *token.Payload) (int, response.Todo) {
	if payload.UserID != body.UserID {
		return http.StatusUnauthorized, response.Todo{}
	}

	result, err := tl.storage.CreateTodo(body)
	log.Println(err)
	if err != nil {
		return http.StatusInternalServerError, response.Todo{}
	}

	return http.StatusOK, result
}

func (tl *TodoLogic) DeleteAllTodo(body request.DeleteAllTodo, payload *token.Payload) (int, string) {
	if payload.UserID != body.UserID {
		return http.StatusUnauthorized, StatusUnauthorized
	}

	err := tl.storage.DeleteAllTodo(body)
	if err != nil {
		return http.StatusInternalServerError, StatusInternalServerError
	}

	return http.StatusOK, "delete successful"
}

func (tl *TodoLogic) GetTodo(todoID string) (int, response.Todo) {
	result, err := tl.storage.GetTodo(todoID)
	if err != nil {
		return http.StatusInternalServerError, response.Todo{}
	}

	return http.StatusOK, result
}

func (tl *TodoLogic) UpdateTodo(body request.UpdateTodo, payload *token.Payload) (int, response.Todo) {
	if payload.UserID != body.UserID {
		return http.StatusUnauthorized, response.Todo{}
	}

	result, err := tl.storage.UpdateTodo(body)
	if err != nil {
		return http.StatusInternalServerError, response.Todo{}
	}

	return http.StatusOK, result
}

func (tl *TodoLogic) DeleteTodo(todoID string) (int, string) {
	err := tl.storage.DeleteTodo(todoID)
	if err != nil {
		return http.StatusInternalServerError, StatusInternalServerError
	}

	return http.StatusOK, "delete successful"
}
