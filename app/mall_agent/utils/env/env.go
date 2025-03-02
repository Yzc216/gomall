package env

import (
	"fmt"
	"os"
)

// MustHasEnvs 检查必要环境变量
func MustHasEnvs(keys ...string) {
	for _, key := range keys {
		if os.Getenv(key) == "" {
			panic(fmt.Sprintf("环境变量 %s 未设置", key))
		}
	}
}

// GetEnvOr 获取环境变量值，如果不存在则返回默认值
func GetEnvOr(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
