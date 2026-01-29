package handlers

import (
	"fiber-template/internal/logs/dto"
	"fiber-template/internal/logs/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// LogHandler handles log API requests
type LogHandler struct {
	svc services.LogService
}

// NewLogHandler creates a new log handler
func NewLogHandler(svc services.LogService) *LogHandler {
	return &LogHandler{svc: svc}
}

// GetLogs returns paginated log history from Postgres, newest first.
// GET /api/v1/logs?page=1&page_size=20
func (h *LogHandler) GetLogs(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "20"))

	list, total, err := h.svc.List(c.Context(), page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	items := make([]*dto.LogItem, 0, len(list))
	for _, m := range list {
		items = append(items, dto.FromApiLog(m))
	}

	return c.JSON(dto.LogListResponse{
		Data:  items,
		Total: total,
		Page:  page,
		Size:  pageSize,
	})
}
