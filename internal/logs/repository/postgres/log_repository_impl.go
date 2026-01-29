package postgres

import (
	"context"
	"errors"

	"fiber-template/internal/logs/models"
	"fiber-template/internal/logs/repository"

	"gorm.io/gorm"
)

var errDBNil = errors.New("database connection is nil")

type logRepository struct {
	db *gorm.DB
}

// NewLogRepository creates a new GORM-based log repository
func NewLogRepository(db *gorm.DB) repository.LogRepository {
	return &logRepository{db: db}
}

func (r *logRepository) Create(ctx context.Context, log *models.ApiLog) error {
	if r == nil || r.db == nil {
		return errDBNil
	}
	return r.db.WithContext(ctx).Create(log).Error
}

func (r *logRepository) List(ctx context.Context, offset, limit int) ([]*models.ApiLog, int64, error) {
	if r == nil || r.db == nil {
		return nil, 0, errDBNil
	}
	var total int64
	if err := r.db.WithContext(ctx).Model(&models.ApiLog{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var list []*models.ApiLog
	err := r.db.WithContext(ctx).
		Model(&models.ApiLog{}).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&list).Error
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
