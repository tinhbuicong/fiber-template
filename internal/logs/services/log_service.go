package services

import (
	"context"

	"fiber-template/internal/logs/models"
)

// LogService defines the log service interface
type LogService interface {
	Save(ctx context.Context, log *models.ApiLog) error
	List(ctx context.Context, page, pageSize int) ([]*models.ApiLog, int64, error)
}
