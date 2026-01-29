package repository

import (
	"context"

	"fiber-template/internal/logs/models"
)

// LogRepository defines the API log repository interface
type LogRepository interface {
	Create(ctx context.Context, log *models.ApiLog) error
	List(ctx context.Context, offset, limit int) ([]*models.ApiLog, int64, error)
}
