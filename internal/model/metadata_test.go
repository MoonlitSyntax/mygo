package model

import (
	"encoding/json"
	"testing"
	"time"
)

func TestCustomTime_MarshalJSON(t *testing.T) {
	// 测试正常序列化
	ct := CustomTime(time.Date(2023, 4, 1, 12, 34, 56, 0, time.UTC))
	data, err := json.Marshal(&ct)
	if err != nil {
		t.Fatalf("MarshalJSON 出错: %v", err)
	}
	expected := `"2023-04-01 12:34:56"`
	if string(data) != expected {
		t.Errorf("序列化结果错误: got %s, want %s", string(data), expected)
	}

	// 测试零值序列化
	ct = CustomTime{}
	data, err = json.Marshal(&ct)
	if err != nil {
		t.Fatalf("MarshalJSON 出错: %v", err)
	}
	if string(data) != "null" {
		t.Errorf("序列化零值错误: got %s, want null", string(data))
	}
}
