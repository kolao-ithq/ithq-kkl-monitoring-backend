package services

import (
	"codelab/go-gin-template/internal/api/middleware/response"
	"codelab/go-gin-template/internal/core/modules/users/domain/dto"
	"codelab/go-gin-template/internal/core/modules/users/domain/entities"
	"codelab/go-gin-template/internal/core/modules/users/domain/models"
	"codelab/go-gin-template/internal/core/modules/users/ports/output"
	"codelab/go-gin-template/internal/database/interfaces"

	"codelab/go-gin-template/internal/core/modules/users/ports/input"
	logs "codelab/go-gin-template/pkg/logger"
	"codelab/go-gin-template/pkg/utils/convertor"

	"codelab/go-gin-template/pkg/constants/keys"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type userService struct {
	redisRepo interfaces.RedisRepository
	userRepo  output.UserRepository
}

var (
	RedisRootKey       = keys.MyApp
	UserExpirationTime = keys.ExpireTime
)

// Implement services
// var _ input.UserService = (*userService)(nil)
func (s *userService) CreateUser(ctx context.Context, user *models.UserRequest) error {
	entities, err := dto.NewUserMapper().UserModelToEntity(user)
	if err != nil {
		return err
	}
	if err := s.userRepo.CreateUser(ctx, entities); err != nil {
		return err
	}

	// Invalidate cache asynchronously (fire-and-forget)
	// This won't block the response and won't fail the operation if cache fails
	// Use background context with timeout to prevent goroutine leaks
	invalidateCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	s.redisRepo.InvalidateEntityCaches(invalidateCtx, keys.UserInvalidationKeys())

	return nil
}

// Paginate users
// Get all users from cache or database
// Cache the data if not found in cache
// Return the data
// If the data is not found in cache, fetch from database
// Cache the data
// Return the data
// If the data is not found in cache, fetch from database
// Cache the data
// Return the data
func (s *userService) GetAllUsers(ctx context.Context, query response.PaginationQuery) (*response.BaseResponse, error) {
	key := fmt.Sprintf(keys.UserPaginationSearchKey, query.Page, query.Search)
	path := RedisRootKey
	logs.Info(fmt.Sprintf("key: %s", key))
	cacheData, err := s.redisRepo.Get(ctx, key, path)
	if err == nil && cacheData != nil && len(cacheData) > 0 {
		var arr []response.BaseResponse
		if err := json.Unmarshal(cacheData, &arr); err != nil {
			logs.Error(fmt.Sprintf("failed to unmarshal cache data: %v", err))
			return nil, err
		}

		if len(arr) == 0 {
			return nil, nil
		}
		pagination := arr[0] // grab the first element
		logs.Info(fmt.Sprintf("from cache: %v", pagination.Success))
		return &pagination, nil
	}

	// 2. Otherwise: DB fetch
	data, dbUsers, err := s.userRepo.GetAllUsers(ctx, query)
	if err != nil {
		return nil, err
	}

	userModels := convertor.ConvertEntitiesToModels[entities.Users, models.UserResponse](dbUsers)
	data.Data = userModels

	// cache the data
	go func() {
		key := fmt.Sprintf(keys.UserPaginationSearchKey, data.Meta.Page, query.Search)
		if err := s.redisRepo.Set(ctx, key, path, data, UserExpirationTime); err != nil {
			logs.Error(fmt.Sprintf("failed to cache data: %v", err))
		} else {
			logs.Info("data cached successfully")
		}
	}()

	logs.Info("from database")
	logs.Info(fmt.Sprintf("from database: %v", data.Success))
	return data, nil
}

func (s *userService) GetUsers(ctx context.Context) ([]models.UserResponse, error) {
	key := keys.UserListKey
	users, err := s.userRepo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	path := RedisRootKey
	userModels := convertor.ConvertEntitiesToModels[entities.Users, models.UserResponse](users)
	// cache the data
	go func() {
		if err := s.redisRepo.Set(ctx, key, path, userModels, UserExpirationTime); err != nil {
			logs.Error(fmt.Sprintf("failed to cache users at $.users: %v", err))
		} else {
			logs.Info("users cached successfully at $.users")
		}
	}()

	return userModels, nil
}

func (s *userService) GetUserByID(ctx context.Context, id uint) (*models.UserResponse, error) {
	key := fmt.Sprintf(keys.UserKey, id)
	path := RedisRootKey
	cacheData, err := s.redisRepo.Get(ctx, key, path)
	if err == nil && cacheData != nil && len(cacheData) > 0 {
		var userModel models.UserResponse
		if err := json.Unmarshal(cacheData, &userModel); err != nil {
			logs.Error(fmt.Sprintf("failed to unmarshal cache data: %v", err))
			return nil, err
		}
		return &userModel, nil
	}
	user, err := s.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	userModel := convertor.ConvertEntityToModel[entities.Users, models.UserResponse](*user)
	key = fmt.Sprintf(keys.UserKey, userModel.ID)
	go func() {
		if err := s.redisRepo.Set(ctx, key, path, userModel, UserExpirationTime); err != nil {
			logs.Error(fmt.Sprintf("failed to cache user at $.user: %v", err))
		} else {
			logs.Info("user cached successfully at $.user")
		}
	}()
	return &userModel, nil
}
func (s *userService) DeleteUser(ctx context.Context, id uint) error {
	err := s.userRepo.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	go func() {
		key := fmt.Sprintf(keys.UserKey, id)
		if err := s.redisRepo.Delete(ctx, key); err != nil {
			logs.Error(fmt.Sprintf("failed to delete user from cache: %v", err))
		}
	}()
	return nil
}

func (s *userService) UpdateUser(ctx context.Context, user *models.UserRequest) error {
	entities := convertor.ConvertModelToEntity[models.UserRequest, entities.Users](*user)
	if err := s.userRepo.UpdateUser(ctx, &entities); err != nil {
		return err
	}
	go func() {
		key := fmt.Sprintf(keys.UserKey, entities.ID)
		if err := s.redisRepo.Delete(ctx, key); err != nil {
			logs.Error(fmt.Sprintf("failed to delete user from cache: %v", err))
		}
	}()
	return nil
}

func NewUserService(userRepo output.UserRepository, redisRepo interfaces.RedisRepository) input.UserService {
	return &userService{
		userRepo:  userRepo,
		redisRepo: redisRepo,
	}
}
