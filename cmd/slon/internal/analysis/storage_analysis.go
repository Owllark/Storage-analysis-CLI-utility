package analysis

import (
	"memory-cli-utility/pkg/filedata"
)

type AnalysisConfig struct {
	Sort       bool
	Percentage bool
}

func StorageAnalysis(path string, config *AnalysisConfig) (Info, error) {
	var res filedata.FileInfo
	var err error
	res, err = filedata.TraverseDirectoryAsync(path)
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
