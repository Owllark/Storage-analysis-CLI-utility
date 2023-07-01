package main

import (
	"flag"
	"fmt"
	"log"
	"memory-cli-utility/utils"
	"os"
)

type analysisConfig struct {
	sort       bool
	percentage bool
	export     bool
	statistics bool
}

type outputConfig struct {
	percentage bool
}

var sortFlag bool
var percentageFlag bool
var exportFlag bool
var statisticsFlag bool

func main() {
	utils.LogoOutput()
	workingDir, _ := os.Getwd()

	flag.BoolVar(&sortFlag, "s", false, "sorting files and directories")
	flag.BoolVar(&percentageFlag, "p", false, "output percent of parent directory size")
	flag.BoolVar(&exportFlag, "e", false, "write results of analyse into file")
	flag.BoolVar(&statisticsFlag, "stat", false, "output statistics of storage using")

	flag.Parse()

	directory := flag.Arg(0)

	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf(" -%s\t\t%s\n", f.Name, f.Usage)
	})

	workingDir += "D:\\"
	info, err := StorageAnalysis(directory, &analysisConfig{
		sort:       sortFlag,
		percentage: percentageFlag,
		export:     exportFlag,
		statistics: statisticsFlag,
	})
	if err != nil {
		log.Fatal(err)
	}
	InfoOutput(&info, &outputConfig{
		percentage: percentageFlag,
	})
}
