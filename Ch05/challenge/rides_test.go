package main

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxRideSpeed(t *testing.T) {

	reader1 := strings.NewReader(mock1)
	reader2 := strings.NewReader(mock2)

	tests := []struct {
		name   string
		reader io.Reader
		output float64
		err    bool
	}{
		{
			name:   "happy case",
			reader: reader1,
			output: 12.825,
			err:    false,
		},
		{
			name:   "case 2",
			reader: reader2,
			output: 40.5,
			err:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := maxRideSpeed(tt.reader)
			if tt.err {
				assert.NotNil(t, err, "error is expected")
			}
			assert.Nil(t, err, "error not expected")
			assert.Equal(t, tt.output, res)
		})
	}
}
