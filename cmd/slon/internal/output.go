package internal

import (
	"fmt"
	"io"
	"memory-cli-utility/utils"
)

type OutputConfig struct {
	Percentage bool
	MaxNesting uint
}

func InfoOutput(info *Info, config *OutputConfig, output io.Writer) {
	infoOutput(info, 0, config, output)
}
func infoOutput(info *Info, nesting int, config *OutputConfig, output io.Writer) {

	if uint(nesting) > config.MaxNesting {
		return
	}
	var res string

	for i := 0; i < nesting; i++ {
		res += "\t"
	}
	res += fmt.Sprintf("%s  -  ", info.Name)
	size := utils.ConvertBytesToHigherValues(info.Size)
	res += GetFileSizeString(size) + "  "
	if config.Percentage {
		res += fmt.Sprintf("%.2f %%", info.Percent)
	}
	fmt.Fprintln(output, res)
	if info.IsDir {
		for i := range info.Children {
			infoOutput(&info.Children[i], nesting+1, config, output)
		}
	}

}

func GetFileSizeString(size utils.FileSize) string {

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
