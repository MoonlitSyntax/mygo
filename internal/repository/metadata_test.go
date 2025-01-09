package repository

import (
	"encoding/json"
	"fmt"
	"mygo/internal/util"
	"os"
	"testing"
)

func TestGetAllArticleMetadata(t *testing.T) {
	_ = os.Chdir("/Users/dionysus/code/GolandProjects/mygo")
	util.InitAll()

	list, _ := GetAllArticleMetadata()

	jsonData, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		t.Fatalf("JSON 格式化失败: %v", err)
	}

	fmt.Println(string(jsonData))
}
