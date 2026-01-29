package services

import (
	"context"

	"fiber-template/internal/logs/models"
	"fiber-template/internal/logs/repository"
)

type logServiceImpl struct {
	repo repository.LogRepository
}

// NewLogService creates a new log service
func NewLogService(repo repository.LogRepository) LogService {
	return &logServiceImpl{repo: repo}
}

func (s *logServiceImpl) Save(ctx context.Context, log *models.ApiLog) error {
	return s.repo.Create(ctx, log)
}

func (s *logServiceImpl) List(ctx context.Context, page, pageSize int) ([]*models.ApiLog, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	return s.repo.List(ctx, offset, pageSize)
}
