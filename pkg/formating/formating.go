package formating

import (
	"errors"
	"fmt"
)

// FileSize stores size of a file as amount of terabytes, gigabytes, megabytes, kilobytes and bytes
type FileSize struct {
	Tb int64
	Gb int64
	Mb int64
	Kb int64
	B  int64
}

// ConvertBytesToHigherValues converts number of bytes to FileSize struct
func ConvertBytesToHigherValues(n int64) (FileSize, error) {
	var res FileSize
	if n < 0 {
		return res, errors.New("cannot convert negative value")
	}
	res.Tb = n >> 40
	n = n & ((1 << 40) - 1)
	res.Gb = n >> 30
	n = n & ((1 << 30) - 1)
	res.Mb = n >> 20
	n = n & ((1 << 20) - 1)
	res.Kb = n >> 10
	n = n & ((1 << 10) - 1)
	res.B = n
	return res, nil
}

// GetFileSizeString returns string representing the given bytes number
// as the biggest unit of data with non-zero integer part
func GetFileSizeString(n int64) string {
	size, err := ConvertBytesToHigherValues(n)
	if err != nil {
		return ""
	}
	if size.Tb > 0 {
		fractionalPart := fmt.Sprintf("%f", float32(size.Gb)/1024)[2:4]
		return fmt.Sprintf("%d.%s Tb", size.Tb, fractionalPart)
	}
	if size.Gb > 0 {
		fractionalPart := fmt.Sprintf("%f", float32(size.Mb)/1024)[2:4]
		return fmt.Sprintf("%d.%s Gb", size.Gb, fractionalPart)
	}
	if size.Mb > 0 {
		fractionalPart := fmt.Sprintf("%f", float32(size.Kb)/1024)[2:4]
		return fmt.Sprintf("%d.%s Mb", size.Mb, fractionalPart)
	}
	if size.Kb > 0 {
		fractionalPart := fmt.Sprintf("%f", float32(size.B)/1024)[2:4]
		return fmt.Sprintf("%d.%s Kb", size.Kb, fractionalPart)
	}
	return fmt.Sprintf("%d bytes", size.B)

}

// FormatToMinimalPrecision returns string representing given float number with given precision as string
// if the number less then 10^(-minPrecision) will be returned string representing number
// up to the first significant digit
func FormatToMinimalPrecision(num float64, minPrecision uint) string {

	var res string

	if num < 0 {
		num *= -1
		res += "-"
	}

	border := float64(1)
	for i := uint(0); i < minPrecision; i++ {
		border /= 10
	}

	if num > border || num == 0 {
		formatString := fmt.Sprintf("%%.%df", minPrecision)
		res += fmt.Sprintf(formatString, num)
	} else {
		precision := minPrecision
		for num < border {
			border /= 10
			precision++
		}
		formatString := fmt.Sprintf("%%.%df", precision)
		res += fmt.Sprintf(formatString, num)
	}
	return res
}
