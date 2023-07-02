package main

import (
	"flag"
	"fmt"
	"log"
	"memory-cli-utility/cmd/slon/internal/analysis"
	"memory-cli-utility/utils"
	"os"
)

var sortFlag bool
var percentageFlag bool
var exportFlag string
var nestingFlag uint

func main() {

	flag.BoolVar(&sortFlag, "s", false, "sorting files and directories")
	flag.BoolVar(&percentageFlag, "p", false, "output percent of parent directory size")
	flag.StringVar(&exportFlag, "e", "", "write results of analyse into file")
	flag.UintVar(&nestingFlag, "n", ^uint(0), "maximal nestingFlag of files")

	// Flags and Args processing

	flag.Parse()

	args := flag.Args()

	var directory string
	var outputFile string

	if len(args) == 0 {
		utils.LogoOutput()
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Printf(" -%s\t\t%s\n", f.Name, f.Usage)
		})
		return
	}

	directory = args[0]

	if len(args) > 1 {
		log.Fatal("Too many arguments")
	}

	// Analysis

	info, err := analysis.StorageAnalysis(directory, &analysis.AnalysisConfig{
		Sort:       sortFlag,
		Percentage: percentageFlag,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Output

	outputConfig := &analysis.OutputConfig{
		Percentage: percentageFlag,
		MaxNesting: nestingFlag,
	}
	analysis.InfoOutput(&info, outputConfig, os.Stdout)

	if exportFlag != "" {
		file, err := os.Create(outputFile)
		if err != nil {
			fmt.Printf("error writing to the file '%s'\n", outputFile)
		}
		analysis.InfoOutput(&info, outputConfig, file)
	}
}
