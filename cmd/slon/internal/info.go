package internal

import (
	"memory-cli-utility/pkg/file_data"
	"sort"
)

type Info struct {
	Name     string
	Size     int64
	IsDir    bool
	Percent  float32
	Children []Info
}

type BySizeDescending []Info

func (arr BySizeDescending) Len() int {
	return len(arr)
}

func (arr BySizeDescending) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr BySizeDescending) Less(i, j int) bool {
	return arr[i].Size < arr[j].Size
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

func NewInfo(fileInfo file_data.FileInfo) Info {
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
	for i := range info.Children {
		info.Children[i].calculatePercent(info.Size)
	}
}

func (info *Info) CalculateSize() int64 {
	if !info.IsDir {
		return info.Size
	}
	size := int64(0)
	for i := range info.Children {
		size += info.Children[i].CalculateSize()
	}
	info.Size = size
	return size
}

func SortBySizeDescending(info *Info) {
	sort.Slice(info.Children, func(i, j int) bool {
		return info.Children[i].Size > info.Children[j].Size
	})
}

func (info *Info) Sort(sortFunc func(info *Info)) {
	sortFunc(info)
	if !info.IsDir {
		return
	}
	for i := range info.Children {
		info.Children[i].Sort(sortFunc)
	}
}
