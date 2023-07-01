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

type FileSize struct {
	Tb int64
	Gb int64
	Mb int64
	Kb int64
	B  int64
}

func ConvertBytesToHigherValues(n int64) FileSize {
	var res FileSize
	res.Tb = n >> 40
	n = n & ((1 << 40) - 1)
	res.Gb = n >> 30
	n = n & ((1 << 30) - 1)
	res.Mb = n >> 20
	n = n & ((1 << 20) - 1)
	res.Kb = n >> 10
	n = n & ((1 << 10) - 1)
	res.B = n
	return res
}
