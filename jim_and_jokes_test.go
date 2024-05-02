package hackerrankgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJimAndJokes(t *testing.T) {
	tests := []struct {
		input  [][]int32
		output int32
	}{
		{
			input:  [][]int32{{10, 25}, {8, 31}},
			output: 1,
		},
		{
			input:  [][]int32{{2, 25}, {2, 25}},
			output: 0,
		},
		{
			input:  [][]int32{{11, 10}, {10, 11}},
			output: 1,
		},
	}
	for _, tt := range tests {
		actual := jimAndJokes(tt.input)

		assert.Equal(t, tt.output, actual)
	}
}
