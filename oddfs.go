package main

import (
	"fmt"

	"bazil.org/fuse/fs"
)

// OddFS is the filesystem
type OddFS struct {
	t *fs.Tree
}

// Root returns the filesystem root
func (f OddFS) Root() (fs.Node, error) {
	return f.t.Root()
}

// NewOddFS returns a new OddFS instance
func NewOddFS() *OddFS {
	tree := &fs.Tree{}

	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("%03d.bin", i)
		tree.Add(name, NewBinaryFile(100))
	}

	return &OddFS{
		t: tree,
	}
}
