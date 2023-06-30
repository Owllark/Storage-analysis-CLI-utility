package main

import (
	"fmt"
	"memory-cli-utility/internal"
)

func InfoOutput(info Info, offset int) {

	for i := 0; i < offset; i++ {
		fmt.Printf("\t")
	}
	fmt.Printf("%s  -  %d bytes  %.2f %%\n", info.Name, info.Size, info.Percent)
	if info.IsDir {
		for _, child := range info.Children {
			InfoOutput(child, offset+1)
		}
	}

}

func StorageAnalysis(path string) {
	var res internal.FileInfo
	var err error
	res, err = internal.TraverseDirectory(path)
	if err != nil {
		return
	}
	info := NewInfo(res)
	info.CalculateSize()
	info.CalculatePercent()
	InfoOutput(info, 0)
}
