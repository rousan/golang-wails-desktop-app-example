package utils

import (
	"errors"
	"fmt"
	"os"
	"path"
)

func UserDownloadsDir() (downloadsPath string, err error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("can't find the user home dir")
	}

	downloadsPath = path.Join(homeDir, "Downloads")
	err = os.MkdirAll(downloadsPath, 0755)
	if err != nil {
		return "", fmt.Errorf("can't mkdirAll downloads path: %v", err)
	}

	return
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}
