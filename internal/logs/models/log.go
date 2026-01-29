package models

import "time"

// ApiLog represents an API request log entry (table: api_logs)
type ApiLog struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Method    string    `gorm:"index" json:"method"`
	Path      string    `json:"path"`
	Status    int       `gorm:"index" json:"status"`
	Latency   int64     `json:"latency"` // milliseconds
	IP        string    `json:"ip"`
	CreatedAt time.Time `gorm:"index" json:"created_at"`
}

// TableName overrides the table name
func (ApiLog) TableName() string {
	return "api_logs"
}
