package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler is a Gin middleware that captures errors and returns a consistent JSON error response.
// Usage:
//   r := gin.New()
//   r.Use(response.ErrorHandler())
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Step 1: Process the request first.
		c.Next()

		// Step 2: Check if any errors were added to the context.
		if len(c.Errors) == 0 {
			return
		}

		// Step 3: Use the last error attached to the context.
		err := c.Errors.Last().Err

		// Step 4: Derive HTTP status and message based on error type.
		status := http.StatusInternalServerError
		message := ErrMsgInternalServerError

		if appErr, ok := err.(AppError); ok {
			// Our structured application error.
			status = appErr.Status
			message = appErr.Message
		} else if err != nil {
			// Fallback to the raw error message for non-AppError.
			message = err.Error()
		}

		// Step 5: Send a consistent JSON error response.
		resp := BaseResponse{
			Success: false,
			Message: message,
		}
		c.JSON(status, resp)
	}
}

// NewError creates a new AppError with the specified code and message.
func NewError(code int, errMsg string) error {
	return AppError{
		Status:  code,
		Message: errMsg,
	}
}

// ErrorBadRequest creates a new bad request error.
func ErrorBadRequest(errorMessage string) error {
	return AppError{
		Status:  http.StatusBadRequest,
		Message: errorMessage,
	}
}

// ErrorUnprocessableEntity creates a new unprocessable entity error.
func ErrorUnprocessableEntity(errorMessage string) error {
	return AppError{
		Status:  http.StatusUnprocessableEntity,
		Message: errorMessage,
	}
}

// NewErrorResponses converts an error to an appropriate HTTP response using the same pattern as ErrorHandler.
func NewErrorResponses(ctx *gin.Context, err error) error {
	if err == nil {
		return nil
	}

	status := http.StatusInternalServerError
	message := ErrMsgInternalServerError

	if appErr, ok := err.(AppError); ok {
		status = appErr.Status
		message = appErr.Message
	} else {
		message = err.Error()
	}

	response := BaseResponse{
		Success: false,
		Message: message,
	}

	ctx.JSON(status, response)
	return nil
}

// NewAppErrorStatusMessage creates an error with the given status code and message.
func NewAppErrorStatusMessage(statusCode int, err error) error {
	return AppError{
		Status:  statusCode,
		Message: err.Error(),
	}
}

// NewErrorMessageResponse returns an error response with the given message.
func NewErrorMessageResponse(ctx *gin.Context, message interface{}) error {
	response := BaseResponse{
		Success: false,
		Message: message.(string),
	}

	ctx.JSON(http.StatusUnprocessableEntity, response)
	return nil
}

// NewErrorErrMsgInternalServerError returns an internal server error response.
func NewErrorErrMsgInternalServerError(ctx *gin.Context) error {
	response := BaseResponse{
		Success: false,
		Message: ErrMsgInternalServerError,
	}

	ctx.JSON(http.StatusInternalServerError, response)
	return nil
}

// NewErrorErrMsgUnauthorized returns an unauthorized error response.
func NewErrorErrMsgUnauthorized(ctx *gin.Context) error {
	response := BaseResponse{
		Success: false,
		Message: ErrMsgUnauthorized,
	}

	ctx.JSON(http.StatusUnauthorized, response)
	return nil
}

// NewErrorErrMsgUnauthorizedErrMsgInvalidToken returns an invalid token error response.
func NewErrorErrMsgUnauthorizedErrMsgInvalidToken(ctx *gin.Context) error {
	response := BaseResponse{
		Success: false,
		Message: ErrMsgInvalidAccessToken,
	}

	ctx.JSON(http.StatusUnauthorized, response)
	return nil
}

// NewErrorBadRequest returns a bad request error response.
func NewErrorBadRequest(ctx *gin.Context) error {
	response := BaseResponse{
		Success: false,
		Message: ErrMsgBadRequest,
	}

	ctx.JSON(http.StatusBadRequest, response)
	return nil
}

// NewErrorIDISRequired returns an error response when an ID is required.
func NewErrorIDISRequired(ctx *gin.Context) error {
	response := BaseResponse{
		Success: false,
		Message: ErrMsgParamIdIsRequired,
	}

	ctx.JSON(http.StatusBadRequest, response)
	return nil
}

// NewErrorUnAuthorizeRole returns an error response for unauthorized role.
func NewErrorUnAuthorizeRole(ctx *gin.Context) error {
	response := BaseResponse{
		Success: false,
		Message: YourRoleNotAllowedToAccessThisResource,
	}

	ctx.JSON(http.StatusForbidden, response)
	return nil
}

// NewErrorUnAuthorizePermission returns an error response for unauthorized permission.
func NewErrorUnAuthorizePermission(ctx *gin.Context) error {
	response := BaseResponse{
		Success: false,
		Message: YourPermissionNotAllowedToAccessThisResource,
	}

	ctx.JSON(http.StatusForbidden, response)
	return nil
}
