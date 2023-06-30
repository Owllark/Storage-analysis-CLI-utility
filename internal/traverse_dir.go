package internal

import (
	"os"
	"path/filepath"
)

func TraverseDirectory(path string) (FileInfo, error) {

	var resFile FileInfo
	fileStat, err := os.Stat(path)
	if err != nil {
		return resFile, err
	}
	resFile = FileInfo{
		Name:  fileStat.Name(),
		Size:  fileStat.Size(),
		IsDir: fileStat.IsDir(),
	}
	if !resFile.IsDir {
		return resFile, nil
	}
	content, err := os.ReadDir(path)
	if err != nil {
		return resFile, err
	}
	for _, file := range content {
		child, err := TraverseDirectory(filepath.Join(path, file.Name()))
		if err != nil {
			continue
		}
		resFile.Children = append(resFile.Children, child)
	}
	return resFile, nil
}
