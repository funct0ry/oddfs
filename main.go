package main

import (
	"errors"
	"fmt"
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

func mount(mp string) error {

	if len(mp) == 0 {
		return errors.New("invalid mount point")
	}

	log.Info("Starting ", os.Args[0][2:], "... mounting ", mp)

	c, err := fuse.Mount(
		mp,
		fuse.FSName("oddfs"),
		fuse.Subtype("oddfs"),
		fuse.LocalVolume(),
		fuse.VolumeName("OddFS Volume"),
	)

	if err != nil {
		return err
	}

	defer c.Close()

	err = fs.Serve(c, NewOddFS())

	if err != nil {
		return err
	}

	<-c.Ready

	if err := c.MountError; err != nil {
		return err
	}

	return nil
}

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "ERROR: must specify mount point")
		os.Exit(1)
	}

	err := mount(flag.Arg(0))

	if err != nil {
		log.Fatal(err)
	}
}
