package main

import (
	"fmt"
	"memory-cli-utility/cmd/slon/internal"
	"memory-cli-utility/pkg/file_data"
)

func InfoOutput(info *internal.Info, config *outputConfig) {
	infoOutput(info, 0, config)
}
func infoOutput(info *internal.Info, offset int, config *outputConfig) {

	var res string

	for i := 0; i < offset; i++ {
		res += "\t"
	}
	res += fmt.Sprintf("%s  -  %d bytes  ", info.Name, info.Size)
	if config.percentage {
		res += fmt.Sprintf("%.2f %%", info.Percent)
	}
	fmt.Println(res)
	if info.IsDir {
		for i := range info.Children {
			infoOutput(&info.Children[i], offset+1, config)
		}
	}

}

func StorageAnalysis(path string, config *analysisConfig) (internal.Info, error) {
	var res file_data.FileInfo
	var err error
	res, err = file_data.TraverseDirectory(path)
	if err != nil {
		return internal.Info{}, err
	}
	info := internal.NewInfo(res)
	info.CalculateSize()

	if config.percentage {
		info.CalculatePercent()
	}
	if config.sort {
		info.Sort(internal.SortBySizeDescending)
	}
	if config.export {

	}

	return info, nil
}
