package main

import (
	"memory-cli-utility/internal"
)

type Info struct {
	Name     string
	Size     int64
	IsDir    bool
	Percent  float32
	Children []Info
}

func (info *Info) GetSize() int64 {
	if info.IsDir {
		size := int64(0)
		for _, child := range info.Children {
			size += child.GetSize()
		}
		return size
	} else {
		return info.Size
	}
}

func NewInfo(fileInfo internal.FileInfo) Info {
	var res = Info{
		Name:     fileInfo.Name,
		Size:     fileInfo.Size,
		IsDir:    fileInfo.IsDir,
		Children: nil,
	}

	if fileInfo.IsDir {
		for _, child := range fileInfo.Children {

			res.Children = append(res.Children, NewInfo(child))
		}
	}

	return res
}

func (info *Info) CalculatePercent() {
	info.calculatePercent(info.Size)
}

func (info *Info) calculatePercent(totalSize int64) {
	info.Percent = 100.0 * float32(info.Size) / float32(totalSize)
	if !info.IsDir {
		return
	}
	for i, _ := range info.Children {
		info.Children[i].calculatePercent(info.Size)
	}
}

func (info *Info) CalculateSize() int64 {
	if !info.IsDir {
		return info.Size
	}
	size := int64(0)
	for i, _ := range info.Children {
		size += info.Children[i].CalculateSize()
	}
	info.Size = size
	return size
}

func (info *Info) Inc() {
	info.Size = 100
}
