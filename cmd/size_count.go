package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

type Info struct {
	Name     string
	Size     int64
	IsDir    bool
	Children []Info
}

func InfoOutput(info Info, offset int) {
	for i := 0; i < offset; i++ {
		fmt.Printf("\t")
	}
	fmt.Printf("%s  -  %d bytes\n", info.Name, info.Size)
	if info.IsDir {
		for _, child := range info.Children {
			InfoOutput(child, offset+1)
		}
	}

}

func CountSize(path string) {
	var res Info
	var err error
	res, err = traverseDirectory(path)
	if err != nil {
		return
	}
	InfoOutput(res, 0)
}

func traverseDirectory(path string) (Info, error) {

	var resFile Info
	fileStat, err := os.Stat(path)
	if err != nil {
		return resFile, err
	}
	resFile = Info{
		Name:     fileStat.Name(),
		Size:     fileStat.Size(),
		IsDir:    fileStat.IsDir(),
		Children: make([]Info, 0),
	}
	if !resFile.IsDir {
		return resFile, nil
	}
	content, err := os.ReadDir(path)
	if err != nil {
		return resFile, err
	}
	for _, file := range content {
		child, err := traverseDirectory(filepath.Join(path, file.Name()))
		if err != nil {
			continue
		}
		resFile.Children = append(resFile.Children, child)
	}
	return resFile, nil
}
