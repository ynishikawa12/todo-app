package model

import "time"

type TaskTag struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
