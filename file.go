package main

import (
	"context"

	"bazil.org/fuse"
)

// File represents a regular text file in OddFS
type File struct {
	inode uint64
	data  []byte
}

// Attr returns standard metadata for the file node
func (f File) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Inode = f.inode
	a.Mode = 0444
	a.Size = uint64(len(f.data))
	return nil
}

// ReadAll returns the file data
func (f File) ReadAll(ctx context.Context) ([]byte, error) {
	return f.data, nil
}

// NewTextFile returns a new text file with hex encoded random
// bytes of a given size
func NewTextFile(size int) *File {
	fdata, _ := randomHexBytes(size)
	return &File{
		inode: nextInode(),
		data:  fdata,
	}
}

// NewBinaryFile returns a new binary file with random
// bytes of a given size
func NewBinaryFile(size int) *File {
	fdata, _ := randomBytes(size)
	return &File{
		inode: nextInode(),
		data:  fdata,
	}
}
