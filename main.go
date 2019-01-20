package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

func main() {
	flag.Parse()
	log.Info("Starting ", os.Args[0][2:], "...")
}
