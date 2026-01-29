package dto

// LogListRequest holds query params for GET /api/v1/logs (optional, defaults in handler)
type LogListRequest struct {
	Page     int `query:"page"`
	PageSize int `query:"page_size"`
}
