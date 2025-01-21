package model

import (
	"database/sql"
	"time"
)

type Task struct {
	ID          int            `json:"id" gorm:"primaryKey"`
	UserID      int            `json:"user_id"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	Status      TaskStatus     `json:"status"`
	Priority    TaskPriority   `json:"priority"`
	DueDate     time.Time      `json:"due_date"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}
