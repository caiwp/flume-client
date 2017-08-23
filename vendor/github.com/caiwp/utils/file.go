package utils

import "os"

func IsFile(filePath string) bool {
	f, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return !f.IsDir()
}
