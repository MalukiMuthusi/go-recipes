package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		name        string
		inputFunc   func(int) bool
		inputValues []int
		output      []int
	}{
		{
			name:        "happy case",
			inputFunc:   isOdd,
			inputValues: []int{1, 2, 3, 4, 5, 6, 7, 8},
			output:      []int{1, 3, 5, 7},
		},
		{
			name:        "when output returns and empty slice",
			inputFunc:   isOdd,
			inputValues: []int{2, 4},
			output:      []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vals := filter(tt.inputFunc, tt.inputValues)
			assert.Equal(t, tt.output, vals)
		})
	}
}
