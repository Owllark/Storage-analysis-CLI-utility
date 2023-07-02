package analysis

import (
	"memory-cli-utility/pkg/filedata"
	"sort"
	"sync"
)

type Info struct {
	Name     string
	Size     int64
	IsDir    bool
	Percent  float64
	Children []Info
}

func NewInfo(fileInfo filedata.FileInfo) Info {
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
	info.Percent = 100.0 * float64(info.Size) / float64(totalSize)
	if !info.IsDir {
		return
	}
	for i := range info.Children {
		info.Children[i].calculatePercent(info.Size)
	}
}

func (info *Info) CalculatePercentAsync() {
	var wg sync.WaitGroup
	wg.Add(1)
	go info.calculatePercentAsync(info.Size, &wg)
	wg.Wait()

}

func (info *Info) calculatePercentAsync(totalSize int64, wg *sync.WaitGroup) {
	defer wg.Done()
	info.Percent = 100.0 * float64(info.Size) / float64(totalSize)
	if !info.IsDir {
		return
	}
	for i := range info.Children {
		wg.Add(1)
		go info.Children[i].calculatePercentAsync(info.Size, wg)
	}
}

func (info *Info) CalculateSize() {
	info.calculateSize()
}

func (info *Info) calculateSize() int64 {
	if !info.IsDir {
		return info.Size
	}
	size := int64(0)
	for i := range info.Children {
		size += info.Children[i].calculateSize()
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