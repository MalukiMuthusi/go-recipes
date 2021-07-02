package main

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCmdFreq(t *testing.T) {
	mockData1 := `
	: 1542898478:0;go help mod
	: 1542898478:0;go help mod
	: 1542898485:0;go install app
	`

	mockData2 := `
	: 1542898478:0;go help mod
	: 1542898478:0;go help mod
	: 1542898485:0;go install app
	: 1542898485:0;go install app
	`

	tests := []struct {
		name        string
		inputReader io.Reader
		output      map[string]int
		err         bool
	}{
		{
			name:        "happy case",
			inputReader: strings.NewReader(mockData1),
			output:      map[string]int{"help": 2, "install": 1},
			err:         false,
		},
		{
			name:        "io reader from string",
			inputReader: strings.NewReader(mockData2),
			output:      map[string]int{"help": 2, "install": 2},
			err:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := cmdFreq(tt.inputReader)
			if tt.err {
				assert.NotNil(t, err)
			}
			assert.Nil(t, err)
			assert.Equal(t, tt.output, res)
		})
	}
}
