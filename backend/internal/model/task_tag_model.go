package model

import "time"

type Tag struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
