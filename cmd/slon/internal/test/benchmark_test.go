package test

import (
	"memory-cli-utility/cmd/slon/internal/analysis"
	"memory-cli-utility/pkg/filedata"
	"testing"
)

var testDirectoryPath = "D:\\"
var testFileData filedata.FileInfo
var testRawInfo analysis.Info
var testCalculatedSizeInfo analysis.Info

func init() {
	var err error

	testFileData, err = filedata.TraverseDirectory(testDirectoryPath)
	if err != nil {
		return
	}

	info := analysis.NewInfo(testFileData)
	testRawInfo = info

	info.CalculateSize()
	testCalculatedSizeInfo = info

}

func BenchmarkTraverseDirectory(b *testing.B) {
	filedata.TraverseDirectory(testDirectoryPath)
}

func BenchmarkTraverseDirectoryAsync(b *testing.B) {
	filedata.TraverseDirectoryAsync(testDirectoryPath)
}

func BenchmarkNewInfo(b *testing.B) {
	analysis.NewInfo(testFileData)
}

func BenchmarkInfo_CalculateSize(b *testing.B) {
	testRawInfo.CalculateSize()
}

func BenchmarkInfo_CalculatePercent(b *testing.B) {
	testCalculatedSizeInfo.CalculatePercent()
}

func BenchmarkInfo_CalculatePercentAsync(b *testing.B) {
	testCalculatedSizeInfo.CalculatePercentAsync()
}

func BenchmarkInfo_Sort(b *testing.B) {
	testCalculatedSizeInfo.Sort(analysis.SortBySizeDescending)
}
