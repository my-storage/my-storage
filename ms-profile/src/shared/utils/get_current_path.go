package utils

import (
	"os"
	"path/filepath"
)

func GetCurrentPath() (*string, error) {
	dir, _ := os.Getwd()

	path, err := filepath.Abs(filepath.Join(dir))
	if err != nil {
		return nil, err
	}

	return &path, nil
}
