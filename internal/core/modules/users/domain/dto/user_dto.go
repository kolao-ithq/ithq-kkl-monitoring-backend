package dto

import (
	"codelab/go-gin-template/internal/core/modules/users/domain/entities"
	"codelab/go-gin-template/internal/core/modules/users/domain/models"
	"codelab/go-gin-template/pkg/utils/datetime"
	"codelab/go-gin-template/pkg/utils/encryption"
	"strings"

	"github.com/kataras/jwt"
	"gorm.io/gorm"
)

// UserMapper provides methods to convert between domain User and DTO objects
type UserMapper struct {
}

// NewUserMapper creates a new UserMapper
func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (m *UserMapper) UserModelToEntity(user *models.UserRequest) (*entities.Users, error) {
	passwordHash, err := encryption.Encrypt(user.Password)
	if err != nil {
		return nil, err
	}
	return &entities.Users{
		Model:        gorm.Model{ID: user.ID},
		Email:        user.Email,
		PasswordHash: string(passwordHash),
	}, nil
}

func (m *UserMapper) UserEntityToModel(user *entities.Users) *models.UserResponse {
	return &models.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: datetime.NewDateTimeFormatToString(user.CreatedAt),
		UpdatedAt: datetime.NewDateTimeFormatToString(user.UpdatedAt),
	}
}

func (m *UserMapper) ModelsToEntities(users []models.UserRequest) []*entities.Users {
	var entities []*entities.Users
	for _, user := range users {
		entity, err := m.UserModelToEntity(&user)
		if err != nil {
			return nil
		}
		entities = append(entities, entity)
	}
	return entities
}

func (m *UserMapper) TokenPairToJWTDTO(jwt *jwt.TokenPair) *models.UserJWT {
	return &models.UserJWT{
		Access:  strings.Replace(string(jwt.AccessToken), "\"", "", -1),
		Refresh: strings.Replace(string(jwt.RefreshToken), "\"", "", -1),
	}
}
