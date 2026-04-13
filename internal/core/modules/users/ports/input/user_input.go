package input

import (
	"codelab/go-gin-template/internal/api/middleware/response"
	"codelab/go-gin-template/internal/core/modules/users/domain/models"
	"context"
)

type UserService interface {
	CreateUser(ctx context.Context, user *models.UserRequest) error
	UpdateUser(ctx context.Context, user *models.UserRequest) error
	GetAllUsers(ctx context.Context, query response.PaginationQuery) (*response.BaseResponse, error)
	DeleteUser(ctx context.Context, id uint) error
	GetUsers(ctx context.Context) ([]models.UserResponse, error)
	GetUserByID(ctx context.Context, id uint) (*models.UserResponse, error)
}
