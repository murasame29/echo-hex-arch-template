package gateway

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/murasame29/echo-hex-arch-template/pkg/internal/helpers/token"

	"github.com/murasame29/echo-hex-arch-template/pkg/core/entities/models"
	"github.com/murasame29/echo-hex-arch-template/pkg/core/entities/request"
	"github.com/murasame29/echo-hex-arch-template/pkg/core/entities/response"
	"github.com/murasame29/echo-hex-arch-template/pkg/core/infrastructure/storage"
	"github.com/murasame29/echo-hex-arch-template/pkg/internal/helpers/password"
)

type UserGateway interface {
	ListUser(req request.ListUser) (int, []response.ListUser)
	CreateUser(req request.CreateUser) (int, response.User)
	DeleteAllUser() (int, string)
	GetUser(userID string, payload *token.Payload) (int, response.User)
	UpdateUser(body request.UpdateUser, payload *token.Payload) (int, response.User)
	DeleteUser(userID string, payload *token.Payload) (int, string)
}

type UserLogic struct {
	storage storage.UserStorage
}

func NewUserGateway(ctx context.Context, db *sql.DB) UserGateway {
	return &UserLogic{storage.NewUserStorage(ctx, db)}
}

func (ul *UserLogic) ListUser(req request.ListUser) (int, []response.ListUser) {
	result, err := ul.storage.ListUser(req)
	if err != nil {
		return http.StatusInternalServerError, nil
	}
	return http.StatusOK, result
}

func (ul *UserLogic) CreateUser(req request.CreateUser) (int, response.User) {

	hashedPassword, err := password.HashPassword(req.Password)
	if err != nil {
		return http.StatusInternalServerError, response.User{}
	}
	req.Password = hashedPassword

	result, err := ul.storage.CreateUser(req)

	if err != nil {
		return http.StatusInternalServerError, response.User{}
	}

	return http.StatusOK, result
}

func (ul *UserLogic) DeleteAllUser() (int, string) {
	err := ul.storage.DeleteAllUser()
	if err != nil {
		return http.StatusInternalServerError, StatusInternalServerError
	}
	return http.StatusOK, "Delete Successful"
}

func (ul *UserLogic) GetUser(userID string, payload *token.Payload) (int, response.User) {
	if payload.UserID != userID {
		return http.StatusUnauthorized, response.User{}
	}

	result, err := ul.storage.GetUser(userID)
	if err != nil {
		return http.StatusInternalServerError, response.User{}
	}
	return http.StatusOK, result
}

func (ul *UserLogic) UpdateUser(req request.UpdateUser, payload *token.Payload) (int, response.User) {
	if payload.UserID != req.UserID {
		return http.StatusUnauthorized, response.User{}
	}

	hashedPassword, err := password.HashPassword(req.Password)
	if err != nil {
		return http.StatusInternalServerError, response.User{}
	}
	req.Password = hashedPassword

	result, err := ul.storage.UpdateUser(models.UpdateUserParam{
		UserID:   req.UserID,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return http.StatusInternalServerError, response.User{}
	}
	return http.StatusOK, result
}

func (ul *UserLogic) DeleteUser(userID string, payload *token.Payload) (int, string) {
	if payload.UserID != userID {
		return http.StatusUnauthorized, ""
	}

	err := ul.storage.DeleteUser(userID)
	if err != nil {
		return http.StatusInternalServerError, StatusInternalServerError
	}
	return http.StatusOK, "Delete Successful"
}
