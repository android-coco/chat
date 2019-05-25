package util

import (
	"os"
	"path/filepath"
)

func GetAbsPath(relativePath string) (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]) + relativePath)
}
