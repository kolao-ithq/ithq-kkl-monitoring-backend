package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewSuccessResponse sends a successful JSON response with data using gin.
func NewSuccessResponse(ctx *gin.Context, data interface{}) error {
	response := BaseResponse{
		Success: true,
		Data:    data,
	}
	ctx.JSON(http.StatusOK, response)
	return nil
}

// NewSuccessMessagesResponse sends a successful JSON response with a message using gin.
func NewSuccessMessagesResponse(ctx *gin.Context, message string) error {
	response := BaseResponse{
		Success: true,
		Message: message,
	}
	ctx.JSON(http.StatusOK, response)
	return nil
}

// NewCreatedResponse sends a 201 Created JSON response with data using gin.
func NewCreatedResponse(ctx *gin.Context, data interface{}) error {
	response := BaseResponse{
		Success: true,
		Data:    data,
	}
	ctx.JSON(http.StatusCreated, response)
	return nil
}

// NewPaginatedResponse sends a paginated JSON response using gin.
func NewPaginatedResponse(ctx *gin.Context, data interface{}, meta Meta) error {
	response := BaseResponse{
		Success: true,
		Data:    data,
		Meta: &Meta{
			Total:     meta.Total,
			Page:      meta.Page,
			Limit:     meta.Limit,
			TotalPage: meta.TotalPage,
		},
	}
	ctx.JSON(http.StatusOK, response)
	return nil
}
