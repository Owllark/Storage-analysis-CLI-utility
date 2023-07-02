package analysis

import (
	"fmt"
	"io"
	"memory-cli-utility/pkg/formating"
)

type OutputConfig struct {
	Percentage bool
	MaxNesting uint
}

// InfoOutput runs recursion for output information to io.Writer according to OutputConfig
func InfoOutput(info *Info, config *OutputConfig, output io.Writer) {
	infoOutput(info, 0, config, output)
}

// infoOutput recursive function for output information to io.Writer according to OutputConfig
func infoOutput(info *Info, nesting int, config *OutputConfig, output io.Writer) {

	if uint(nesting) > config.MaxNesting {
		return
	}
	var res string

	for i := 0; i < nesting; i++ {
		res += "\t"
	}
	res += fmt.Sprintf("%s  -  ", info.Name)
	res += formating.GetFileSizeString(info.Size) + "  "
	if config.Percentage {
		res += fmt.Sprintf("%s%%", formating.FormatToMinimalPrecision(info.Percent, 2))
	}
	fmt.Fprintln(output, res)
	if info.IsDir {
		for i := range info.Children {
			infoOutput(&info.Children[i], nesting+1, config, output)
		}
	}

}
