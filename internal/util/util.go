package util

import (
	"fmt"
	"mygo/internal/model"
	"strconv"
	"time"
)

func StringToCustomTime(timeStr string) (model.CustomTime, error) {
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return model.CustomTime{}, fmt.Errorf("invalid time format: %w", err)
	}
	return model.CustomTime(parsedTime), nil
}

func StringToUint(s string) (uint, error) {
	u, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return uint(u), nil
}
