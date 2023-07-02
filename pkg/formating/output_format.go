package formating

import "fmt"

type FileSize struct {
	Tb int64
	Gb int64
	Mb int64
	Kb int64
	B  int64
}

func ConvertBytesToHigherValues(n int64) FileSize {
	var res FileSize
	res.Tb = n >> 40
	n = n & ((1 << 40) - 1)
	res.Gb = n >> 30
	n = n & ((1 << 30) - 1)
	res.Mb = n >> 20
	n = n & ((1 << 20) - 1)
	res.Kb = n >> 10
	n = n & ((1 << 10) - 1)
	res.B = n
	return res
}

func GetFileSizeString(n int64) string {
	size := ConvertBytesToHigherValues(n)
	if size.Tb > 0 {
		return fmt.Sprintf("%d.%d Tb", size.Tb, int(1000*float32(size.Gb)/1024))
	}
	if size.Gb > 0 {
		return fmt.Sprintf("%d.%d Gb", size.Gb, int(1000*float32(size.Mb)/1024))
	}
	if size.Mb > 0 {
		return fmt.Sprintf("%d.%d Mb", size.Mb, int(1000*float32(size.Kb)/1024))
	}
	if size.Kb > 0 {
		return fmt.Sprintf("%d.%d Kb", size.Kb, int(1000*float32(size.B)/1024))
	}
	return fmt.Sprintf("%d bytes", size.B)

}

func FormatToMinimalPrecision(num float64, minPrecision int) string {

	var res string

	border := float64(1)
	for i := 0; i < minPrecision; i++ {
		border /= 10
	}

	if num > border || num == 0 {
		formatString := fmt.Sprintf("%%.%df", minPrecision)
		res = fmt.Sprintf(formatString, num)
	} else {
		precision := minPrecision
		for num < border {
			border /= 10
			precision++
		}
		formatString := fmt.Sprintf("%%.%df", precision)
		res = fmt.Sprintf(formatString, num)
	}
	return res
}
