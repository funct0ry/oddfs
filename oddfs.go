package main

import (
	"bazil.org/fuse/fs"
)

// OddFS is the filesystem
type OddFS struct {
	t   *fs.Tree
	cfg *Config
}

// Root returns the filesystem root
func (f OddFS) Root() (fs.Node, error) {
	return f.t.Root()
}

// NewOddFS returns a new OddFS instance
func NewOddFS(c *Config) *OddFS {
	tree := &fs.Tree{}

	for i := 0; i < c.Count; i++ {
		tree.Add(c.Name(i), newFileByConfig(c))
	}

	return &OddFS{
		t:   tree,
		cfg: c,
	}
}

func newFileByConfig(c *Config) *File {
	switch {
	case c.Type == "txt":
		return NewTextFile(c.Size)
	case c.Type == "bin":
		return NewBinaryFile(c.Size)
	default:
		return NewBinaryFile(c.Size)
	}
}
