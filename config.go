package main

import "fmt"

// Config contains all configurations for the OddFS file read from flags
type Config struct {
	Count int
	Type  string
	Size  int
}

// Name constructs and returns a file name from type and number
func (c *Config) Name(n int) string {
	return fmt.Sprintf("%03d.%s", n, c.Type)
}
