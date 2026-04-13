package response

// AppError represents a structured API error
type AppError struct {
	Status  int
	Message string
}

func (a AppError) Error() string {
	return a.Message
}

type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
}

// Meta represents pagination metadata
type Meta struct {
	Total     int `json:"total"`
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalPage int `json:"total_page"`
}
type PaginationQuery struct {
	Page        int    `json:"page,omitempty" query:"page"`
	Limit       int    `json:"limit,omitempty" query:"limit"`
	Search     string `json:"search,omitempty" query:"search"`
	Status      string `json:"status,omitempty" query:"status"`
	Sort        string `json:"sort,omitempty" query:"sort"`
	Type        string `json:"type,omitempty" query:"type"`
}

// Common error messages
const (
	ErrMsgInternalServerError                    = "Internal server error"
	ErrMsgBadRequest                             = "Bad request"
	ErrMsgUnauthorized                           = "Unauthorized"
	ErrMsgInvalidAccessToken                     = "Invalid access token"
	ErrMsgParamIdIsRequired                      = "ID parameter is required"
	YourRoleNotAllowedToAccessThisResource       = "Your role is not allowed to access this resource"
	YourPermissionNotAllowedToAccessThisResource = "Your permission is not allowed to access this resource"
)
