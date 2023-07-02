package filedata

import (
	"errors"
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

func TraverseDirectoryAsync(path string) (FileInfo, error) {

	_, err := os.Stat(path)
	if err != nil {
		return FileInfo{}, errors.New("The system cannot find the file specified.")
	}

	var fileChan = make(chan FileInfo)

	go traverseDirectoryAsync(path, fileChan)
	res := <-fileChan

	return res, nil
}
func traverseDirectoryAsync(path string, fileChan chan FileInfo) {

	defer close(fileChan)
	var resFile FileInfo
	fileStat, err := os.Stat(path)
	if err != nil {
		return
	}
	resFile = FileInfo{
		Name:  fileStat.Name(),
		Size:  fileStat.Size(),
		IsDir: fileStat.IsDir(),
	}
	if !resFile.IsDir {
		fileChan <- resFile
		return
	}
	content, err := os.ReadDir(path)
	if err != nil {
		fileChan <- resFile
		return
	}

	childFileChan := make([]chan FileInfo, len(content))
	for i, file := range content {
		childFileChan[i] = make(chan FileInfo)
		go traverseDirectoryAsync(filepath.Join(path, file.Name()), childFileChan[i])
	}

	for i := range childFileChan {
		file := <-childFileChan[i]
		resFile.Children = append(resFile.Children, file)
	}

	fileChan <- resFile
	return

}
