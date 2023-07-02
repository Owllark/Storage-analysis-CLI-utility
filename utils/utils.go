package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func LogoOutput() {
	dirPath, err := GetPathToUtilityDir()
	if err != nil {
		return
	}
	text, err := os.ReadFile(filepath.Join(dirPath, "utils", "logo.txt"))
	if err != nil {
		return
	}
	fmt.Println(string(text))
}

func GetPathToUtilityDir() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	path, _ := filepath.Split(exePath)
	return path, err
}
