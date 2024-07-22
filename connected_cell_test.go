package hackerrankgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectedCell(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int32
		want  int32
	}{
		{
			name:  "ConnectedCell",
			input: [][]int32{{1, 1, 0}, {1, 0, 0}, {0, 0, 1}},
			want:  3,
		},
		{
			name:  "ConnectedCell",
			input: [][]int32{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 1, 0}, {1, 0, 0, 0}},
			want:  5,
		},
	}
	for _, tt := range tests {
		actual := ConnectedCell(tt.input)

		assert.Equal(t, tt.want, actual)
	}
}
