package main

import (
	"flag"
	"fmt"
	"memory-cli-utility/utils"
	"os"
)

var sort bool
var percentage bool
var export bool
var statistics bool

func main() {
	utils.LogoOutput()
	workingDir, _ := os.Getwd()

	flag.BoolVar(&sort, "s", false, "sort files in descending order")
	flag.BoolVar(&percentage, "p", false, "help message for flag name")
	flag.BoolVar(&export, "e", false, "help message for flag name")
	flag.BoolVar(&statistics, "stat", false, "help message for flag name")

	flag.Parse()

	directory := flag.Arg(0)

	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf(" -%s\t\t%s", f.Name, f.Usage)
		fmt.Println("------------------")
	})

	workingDir += "D:\\"
	StorageAnalysis(directory)

}
