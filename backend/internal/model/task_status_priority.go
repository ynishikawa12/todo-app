package model

import (
	"database/sql/driver"
	"fmt"
)

type TaskPriority string

const (
	Low    TaskPriority = "low"
	Medium TaskPriority = "medium"
	High   TaskPriority = "high"
)

func (ts *TaskPriority) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		*ts = TaskPriority(v)
		return nil
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
}

func (ts TaskPriority) Value() (driver.Value, error) {
	return string(ts), nil
}
