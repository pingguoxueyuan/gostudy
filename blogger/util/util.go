package util

import (
	"os"
	"path/filepath"
)

func GetRootDir() (rootPath string) {
	exePath := os.Args[0]
	rootPath = filepath.Dir(exePath)
	return rootPath
}
