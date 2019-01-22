package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

var config Config
var helpRequested bool

func init() {
	flag.IntVarP(&config.Count, "count", "n", 10, "number of files to generate.")
	flag.StringVarP(&config.Type, "type", "t", "txt", "type of files to generate. Supported formats are [txt, bin, png]")
	flag.IntVarP(&config.Size, "size", "s", 100, "size in bytes of files to generate.")

	flag.BoolVarP(&helpRequested, "help", "h", false, "Display usage information (this message)")
}

func main() {
	flag.Parse()

	if helpRequested {
		flag.Usage()
		os.Exit(1)
	}

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "ERROR: must specify mount point")
		os.Exit(1)
	}

	err := mount(flag.Arg(0), &config)

	if err != nil {
		log.Fatal(err)
	}
}
