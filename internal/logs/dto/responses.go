package dto

import "fiber-template/internal/logs/models"

// LogListResponse is the paginated response for GET /api/v1/logs
type LogListResponse struct {
	Data  []*LogItem `json:"data"`
	Total int64      `json:"total"`
	Page  int        `json:"page"`
	Size  int        `json:"size"`
}

// LogItem represents a single log entry in the API response
type LogItem struct {
	ID        uint   `json:"id"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	Status    int    `json:"status"`
	Latency   int64  `json:"latency"`
	IP        string `json:"ip"`
	CreatedAt string `json:"created_at"`
}

// FromApiLog converts a models.ApiLog to LogItem
func FromApiLog(m *models.ApiLog) *LogItem {
	if m == nil {
		return nil
	}
	return &LogItem{
		ID:        m.ID,
		Method:    m.Method,
		Path:      m.Path,
		Status:    m.Status,
		Latency:   m.Latency,
		IP:        m.IP,
		CreatedAt: m.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}
