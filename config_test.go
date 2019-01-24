package main

import "testing"

func TestName(t *testing.T) {
	tt := []struct {
		name  string
		ftype string
		fnum  int
		fname string
	}{
		{"binary file", "bin", 1, "001.bin"},
		{"text file", "txt", 1, "001.txt"},
		{"text file", "txt", 100, "100.txt"},
		{"text file", "txt", 1000, "1000.txt"},
		{"image (png) file", "png", 1, "001.png"},
		{"image (jpg) file", "jpg", 1, "001.jpg"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := Config{
				Count: 10,
				Type:  tc.ftype,
				Size:  100,
			}
			n := c.Name(tc.fnum)
			if n != tc.fname {
				t.Fatalf("expected file name for %v to be %v; got %v", tc.name, tc.fname, n)
			}
		})
	}
}
