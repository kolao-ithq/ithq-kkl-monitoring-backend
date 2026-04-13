package repositories

import (
	"codelab/go-gin-template/internal/api/middleware/response"
	"codelab/go-gin-template/internal/core/modules/users/domain/entities"
	"codelab/go-gin-template/internal/core/modules/users/ports/output"
	"codelab/go-gin-template/pkg/utils/calculate"
	"context"

	"gorm.io/gorm"
)

// type UserRepository interface {
// 	CreateUser(ctx context.Context, user *entities.Users) error
// 	GetUserByID(ctx context.Context, id uint) (*entities.Users, error)
// 	GetAllUsers(ctx context.Context, query response.PaginationQuery) (*response.BaseResponse, []entities.Users, error)
// 	GetUsers(ctx context.Context) ([]entities.Users, error)
// 	DeleteUser(ctx context.Context, id uint) error
// 	UpdateUser(ctx context.Context, user *entities.Users) error
// }

type UserRepository struct {
	db *gorm.DB
}

// implement repositories
func (r *UserRepository) CreateUser(ctx context.Context, user *entities.Users) error {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id uint) (*entities.Users, error) {
	var user entities.Users
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAllUsers(ctx context.Context, query response.PaginationQuery) (*response.BaseResponse, []entities.Users, error) {

	var users []entities.Users
	var totalRows int64
	noPagination := false

	// Normalize pagination
	if query.Page < 1 {
		query.Page = 1
	}

	if query.Limit == 0 {
		query.Limit = 10
	}

	if query.Limit < 0 {
		noPagination = true
	}

	// Base query
	baseQuery := r.db.WithContext(ctx).
		Model(&entities.Users{}).
		Select("id", "full_name", "email", "created_at", "updated_at").
		Where("deleted_at IS NULL")

	// Search
	if query.Search != "" {
		query.Page = 1
		baseQuery = baseQuery.Where(
			"full_name ILIKE ? OR email ILIKE ?",
			"%"+query.Search+"%",
			"%"+query.Search+"%",
		)
	}

	// Count total rows (NO offset/limit here)
	if err := baseQuery.Count(&totalRows).Error; err != nil {
		return nil, nil, err
	}

	// Pagination metadata
	totalPage := calculate.CalculatePageSize(totalRows, query.Limit)

	// Apply pagination only if needed
	if !noPagination {
		offset := calculate.CalculateOffset(query.Page, query.Limit)
		baseQuery = baseQuery.Offset(offset).Limit(query.Limit)
	}

	if query.Page > totalPage {
		query.Page = 1
	}

	// Fetch data
	if err := baseQuery.
		Order("created_at DESC").
		Find(&users).Error; err != nil {
		return nil, nil, err
	}

	queryData := &response.BaseResponse{
		Success: true,
		Message: "Success",
		Data:    users,
		Meta: &response.Meta{
			Total:     int(totalRows),
			Page:      query.Page,
			Limit:     query.Limit,
			TotalPage: totalPage,
		},
	}

	return queryData, users, nil
}

func (r *UserRepository) GetUsers(ctx context.Context) ([]entities.Users, error) {
	var users []entities.Users
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id uint) error {
	// use transaction to delete the user
	res := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&entities.Users{}, id).Error; err != nil {
			return err
		}
		return nil
	})
	return res
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *entities.Users) error {
	if err := r.db.WithContext(ctx).Model(&entities.Users{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func NewUserRepository(db *gorm.DB) output.UserRepository {
	return &UserRepository{
		db: db,
	}
}
