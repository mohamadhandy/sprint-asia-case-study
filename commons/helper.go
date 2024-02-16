package commons

import (
	"path/filepath"
	"runtime"
	"strings"
)

func DynamicDir() string {
	_, b, _, _ := runtime.Caller(0)
	bStr := filepath.Dir(b)
	baseDir := strings.Replace(bStr, "/commons", "", -1)

	return baseDir
}
