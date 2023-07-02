package filedata

import (
	"errors"
	"os"
	"path/filepath"
)

// TraverseDirectory gets filepath and runs concurrent recursion for going through the all
// nested files, forming FileInfo result.
func TraverseDirectory(path string) (FileInfo, error) {

	_, err := os.Stat(path)
	if err != nil {
		return FileInfo{}, errors.New("The system cannot find the file specified.")
	}

	var fileChan = make(chan FileInfo)

	go traverseDirectory(path, fileChan)
	res := <-fileChan

	return res, nil
}

// traverseDirectory gets filepath and channel for sending results to upper recursion level
func traverseDirectory(path string, fileChan chan FileInfo) {

	defer close(fileChan)
	var resFile FileInfo
	fileStat, err := os.Stat(path)
	if err != nil {
		return
	}
	resFile = FileInfo{
		Name:    fileStat.Name(),
		Size:    fileStat.Size(),
		IsDir:   fileStat.IsDir(),
		Mode:    fileStat.Mode(),
		ModTime: fileStat.ModTime(),
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
		go traverseDirectory(filepath.Join(path, file.Name()), childFileChan[i])
	}

	for i := range childFileChan {
		file := <-childFileChan[i]
		resFile.Children = append(resFile.Children, file)
	}

	fileChan <- resFile
	return

}

// TraverseDirectorySync gets filepath and recursively goes through the all
// nested files, forming FileInfo result. Works without concurrency
func TraverseDirectorySync(path string) (FileInfo, error) {

	var resFile FileInfo
	fileStat, err := os.Stat(path)
	if err != nil {
		return resFile, err
	}
	resFile = FileInfo{
		Name:    fileStat.Name(),
		Size:    fileStat.Size(),
		IsDir:   fileStat.IsDir(),
		Mode:    fileStat.Mode(),
		ModTime: fileStat.ModTime(),
	}
	if !resFile.IsDir {
		return resFile, nil
	}
	content, err := os.ReadDir(path)
	if err != nil {
		return resFile, err
	}
	for _, file := range content {
		child, err := TraverseDirectorySync(filepath.Join(path, file.Name()))
		if err != nil {
			continue
		}
		resFile.Children = append(resFile.Children, child)
	}
	return resFile, nil
}
