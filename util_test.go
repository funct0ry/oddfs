package main

import (
	"testing"
)

func TestRandomHexBytes(t *testing.T) {

	tests := []struct {
		name    string
		size    uint64
		wantErr bool
	}{
		{"random hex bytes", 0, false},
		{"random hex bytes", 1, false},
		{"random hex bytes", 2, false},
		{"random hex bytes", 3, false},
		{"random hex bytes", 4, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := randomHexBytes(tt.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("%v error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			got := uint64(len(b))
			if got != tt.size {
				t.Errorf("expected size of %v to be %v; got %v", tt.name, tt.size, got)
			}
		})
	}
}
