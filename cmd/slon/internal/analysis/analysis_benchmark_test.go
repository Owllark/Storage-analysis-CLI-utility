package analysis

import (
	"memory-cli-utility/pkg/filedata"
	"testing"
)

var testDirectoryPath = "D:\\"
var testFileData filedata.FileInfo
var testRawInfo Info
var testCalculatedSizeInfo Info

func init() {
	var err error

	testFileData, err = filedata.TraverseDirectorySync(testDirectoryPath)
	if err != nil {
		return
	}

	info := NewInfo(testFileData)
	testRawInfo = info

	info.CalculateSize()
	testCalculatedSizeInfo = info

}

func BenchmarkTraverseDirectory(b *testing.B) {
	filedata.TraverseDirectorySync(testDirectoryPath)
}

func BenchmarkTraverseDirectoryAsync(b *testing.B) {
	filedata.TraverseDirectory(testDirectoryPath)
}

func BenchmarkNewInfo(b *testing.B) {
	NewInfo(testFileData)
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
	testCalculatedSizeInfo.Sort(SortBySizeDescending)
}
