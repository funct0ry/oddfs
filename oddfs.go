package main

import "bazil.org/fuse/fs"

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
	return &OddFS{
		t: &fs.Tree{},
	}
}
