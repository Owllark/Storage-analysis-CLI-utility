package formating

import "testing"

func TestConvertBytesToHigherValues(t *testing.T) {
	var tests = []struct {
		input    int64
		expected FileSize
	}{
		{-1, FileSize{0, 0, 0, 0, 0}},
		{0, FileSize{0, 0, 0, 0, 0}},
		{100, FileSize{0, 0, 0, 0, 100}},
		{1024, FileSize{0, 0, 0, 1, 0}},
		{1048, FileSize{0, 0, 0, 1, 24}},
		{1 << 20, FileSize{0, 0, 1, 0, 0}},
		{1<<20 + 1<<18, FileSize{0, 0, 1, 256, 0}},
		{1 << 30, FileSize{0, 1, 0, 0, 0}},
		{1<<30 + 1<<25, FileSize{0, 1, 32, 0, 0}},
		{1 << 40, FileSize{1, 0, 0, 0, 0}},
		{1<<40 + 1<<39, FileSize{1, 512, 0, 0, 0}},
		{20<<40 + 1<<39, FileSize{20, 512, 0, 0, 0}},
	}

	for _, test := range tests {
		if got, _ := ConvertBytesToHigherValues(test.input); got != test.expected {
			t.Errorf("ConvertBytesToHigherValues(%d) = %v", test.input, got)
		}
	}
}
func TestGetFileSizeString(t *testing.T) {
	var tests = []struct {
		input    int64
		expected string
	}{
		{-1, ""},
		{0, "0 bytes"},
		{100, "100 bytes"},
		{1024, "1.00 Kb"},
		{1048, "1.02 Kb"},
		{1 << 20, "1.00 Mb"},
		{1<<20 + 1<<18, "1.25 Mb"},
		{1 << 30, "1.00 Gb"},
		{1<<30 + 1<<25, "1.03 Gb"},
		{1 << 40, "1.00 Tb"},
		{1<<40 + 1<<39, "1.50 Tb"},
		{2<<40 + 1<<39, "2.50 Tb"},
	}

	for i, test := range tests {
		if got := GetFileSizeString(test.input); got != test.expected {
			t.Errorf("%d. GetFileSizeString(%d) = %v", i, test.input, got)
		}
	}

}

func TestFormatToMinimalPrecision(t *testing.T) {
	var tests = []struct {
		num      float64
		prec     uint
		expected string
	}{
		{-1.001, 2, "-1.00"},
		{0, 5, "0.00000"},
		{100, 2, "100.00"},
		{1000, 1, "1000.0"},
		{0.45, 1, "0.5"},
		{0.0027, 5, "0.00270"},
		{0.00013, 1, "0.0001"},
	}

	for i, test := range tests {
		if got := FormatToMinimalPrecision(test.num, test.prec); got != test.expected {
			t.Errorf("%d. FormatToMinimalPrecision(%f, %d) = %v", i, test.num, test.prec, got)
		}
	}
}
