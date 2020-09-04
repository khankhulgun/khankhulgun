package config

import (
	"path/filepath"
	"runtime"
	"strings"
)

func AbsolutePath() string{
	_, fileName, _, _ := runtime.Caller(0)
	return strings.Trim(filepath.Dir(fileName), "utils")
}