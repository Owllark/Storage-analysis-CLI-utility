package internal

import (
	"memory-cli-utility/pkg/file_data"
)

type AnalysisConfig struct {
	Sort       bool
	Percentage bool
	Statistics bool
}

func StorageAnalysis(path string, config *AnalysisConfig) (Info, error) {
	var res file_data.FileInfo
	var err error
	res, err = file_data.TraverseDirectory(path)
	if err != nil {
		return Info{}, err
	}
	info := NewInfo(res)
	info.CalculateSize()

	if config.Percentage {
		info.CalculatePercent()
	}
	if config.Sort {
		info.Sort(SortBySizeDescending)
	}

	return info, nil
}
