package util

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func PkgPath(day int) (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", fmt.Errorf("failed to generate package path")
	}

	return filepath.Join(filepath.Dir(filepath.Dir(filename)), fmt.Sprintf("day%d", day)), nil
}
