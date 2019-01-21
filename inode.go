package main

import (
	"sync/atomic"
)

var inode uint64

func init() {
	inode = 1
}

func nextInode() uint64 {
	return atomic.AddUint64(&inode, 1)
}
