package model

import (
	"database/sql/driver"
	"fmt"
)

type TaskStatus string

const (
	Pending    TaskStatus = "pending"
	InProgress TaskStatus = "in_progress"
	Completed  TaskStatus = "completed"
)

func (ts *TaskStatus) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		*ts = TaskStatus(v)
		return nil
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
}

func (ts TaskStatus) Value() (driver.Value, error) {
	return string(ts), nil
}
