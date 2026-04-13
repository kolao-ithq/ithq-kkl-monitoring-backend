package handler

import (
	"codelab/go-gin-template/internal/api/middleware/response"
	"codelab/go-gin-template/internal/core/modules/users/domain/models"
	"codelab/go-gin-template/internal/core/modules/users/ports/input"

	"codelab/go-gin-template/pkg/utils/convertor"

	"github.com/gin-gonic/gin"
)

type (
	UserHandler interface {

		// User
		GetUser(ctx *gin.Context) error
		GetUserByID(ctx *gin.Context) error
		UpdateUser(ctx *gin.Context) error
		DeleteUser(ctx *gin.Context) error
		// UploadAvatar(ctx *gin.Context) error
	}
	userHandler struct {
		userService input.UserService
	}
)

// DeleteUser implements UserHandler.
func (h *userHandler) DeleteUser(ctx *gin.Context) error {
	id, err := convertor.StringToUint(ctx.Request.Header.Get("id"))
	if err != nil {
		return response.NewErrorResponses(ctx, err)
	}
	err = h.userService.DeleteUser(ctx.Request.Context(), id)
	if err != nil {
		return response.NewErrorResponses(ctx, err)
	}
	return response.NewSuccessMessagesResponse(ctx, "User deleted successfully")
}

// UpdateUser implements UserHandler.
func (h *userHandler) UpdateUser(ctx *gin.Context) error {
	req := &models.UserRequest{}
	if err := ctx.BindJSON(req); err != nil {
		return response.NewErrorResponses(ctx, err)
	}
	err := h.userService.UpdateUser(ctx.Request.Context(), req)
	if err != nil {
		return response.NewErrorResponses(ctx, err)
	}
	return response.NewSuccessMessagesResponse(ctx, "User updated successfully")
}

// GetUserByID implements UserHandler.
func (h *userHandler) GetUserByID(ctx *gin.Context) error {
	id, err := convertor.StringToUint(ctx.Request.Header.Get("id"))
	if err != nil {
		return response.NewErrorResponses(ctx, err)
	}
	res, err := h.userService.GetUserByID(ctx.Request.Context(), id)
	if err != nil {
		return response.NewErrorResponses(ctx, err)
	}
	return response.NewSuccessResponse(ctx, res)
}

func (h *userHandler) GetUser(c *gin.Context) error {
	res, err := h.userService.GetUsers(c.Request.Context())
	if err != nil {
		return response.NewErrorResponses(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(
	userService input.UserService,
) UserHandler {
	return &userHandler{
		userService: userService,
	}
}
