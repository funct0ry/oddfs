package main

import (
	"errors"
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	log "github.com/sirupsen/logrus"
)

func mount(mp string, cfg *Config) error {

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

	err = fs.Serve(c, NewOddFS(cfg))

	if err != nil {
		return err
	}

	<-c.Ready

	if err := c.MountError; err != nil {
		return err
	}

	return nil
}
