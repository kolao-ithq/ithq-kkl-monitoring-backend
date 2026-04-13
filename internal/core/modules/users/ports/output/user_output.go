package output

import (
	"codelab/go-gin-template/internal/api/middleware/response"
	"codelab/go-gin-template/internal/core/modules/users/domain/entities"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entities.Users) error
	GetUserByID(ctx context.Context, id uint) (*entities.Users, error)
	GetAllUsers(ctx context.Context, query response.PaginationQuery) (*response.BaseResponse, []entities.Users, error)
	GetUsers(ctx context.Context) ([]entities.Users, error)
	DeleteUser(ctx context.Context, id uint) error
	UpdateUser(ctx context.Context, user *entities.Users) error
}
