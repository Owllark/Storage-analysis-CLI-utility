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

	for i := 0; i < b.N; i++ {
		filedata.TraverseDirectorySync(testDirectoryPath)
	}
}

func BenchmarkTraverseDirectoryAsync(b *testing.B) {
	for i := 0; i < b.N; i++ {
		filedata.TraverseDirectory(testDirectoryPath)
	}

}

func BenchmarkNewInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewInfo(testFileData)
	}
}

func BenchmarkInfo_CalculateSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testRawInfo.CalculateSize()
	}

}

func BenchmarkInfo_CalculatePercent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testCalculatedSizeInfo.CalculatePercent()
	}

}

func BenchmarkInfo_CalculatePercentAsync(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testCalculatedSizeInfo.CalculatePercentAsync()
	}

}

func BenchmarkInfo_Sort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testCalculatedSizeInfo.Sort(SortBySizeDescending)
	}

}
