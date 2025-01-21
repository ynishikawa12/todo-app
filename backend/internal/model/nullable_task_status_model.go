package model

import (
	"database/sql/driver"
	"fmt"
)

type NullableTaskStatus struct {
	Status TaskStatus
	Valid  bool
}

func (nts *NullableTaskStatus) Scan(value interface{}) error {
	if value == nil {
		nts.Status, nts.Valid = "", false
		return nil
	}
	nts.Valid = true
	switch v := value.(type) {
	case string:
		nts.Status = TaskStatus(v)
		return nil
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
}

func (nts NullableTaskStatus) Value() (driver.Value, error) {
	if !nts.Valid {
		return nil, nil
	}
	return string(nts.Status), nil
}
