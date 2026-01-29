package logs

import "time"

type ApiLog struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Method    string    `json:"method"`
	Path      string    `json:"path"`
	Status    int       `json:"status"`
	Latency   string    `json:"latency"`
	IP        string    `json:"ip"`
	CreatedAt time.Time `json:"created_at"`
}
