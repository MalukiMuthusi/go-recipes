package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTsConvert(t *testing.T) {
	tests := []struct {
		name   string
		ts     string
		from   string
		to     string
		err    bool
		output string
	}{
		{
			name:   "happy case",
			ts:     "2021-03-08T19:12",
			from:   "America/Los_Angeles",
			to:     "Asia/Jerusalem",
			err:    false,
			output: "2021-03-09T05:12",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			toTime, err := tsConvert(tt.ts, tt.from, tt.to)
			if tt.err {
				assert.NotNil(t, err, "error expected")
			}
			assert.Nil(t, err, "error not expected")
			assert.Equal(t, tt.output, toTime)
		})
	}
}
