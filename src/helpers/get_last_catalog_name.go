package helpers

import (
	"os"
	"path/filepath"
)

func GetLastCatalogName() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	_, last := filepath.Split(pwd)
	return last, err
}
