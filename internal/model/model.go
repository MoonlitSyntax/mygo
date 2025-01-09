package model

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type CustomTime time.Time

type GormModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt CustomTime     `json:"created_at"`
	UpdatedAt CustomTime     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	if ct == nil || time.Time(*ct).IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", time.Time(*ct).Format("2006-01-02 15:04:05"))), nil
}

func (ct *CustomTime) Scan(value interface{}) error {
	if value, ok := value.(time.Time); ok {
		*ct = CustomTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", value)
}

func (ct CustomTime) Value() (driver.Value, error) {
	t := time.Time(ct)
	if t.IsZero() {
		return nil, nil
	}
	return t, nil
}
