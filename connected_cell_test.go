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

func TestAdjacentFilledCell(t *testing.T) {
	type adjacentInput struct {
		grid  Grid
		start point
	}

	tests := []struct {
		expected []point
		input    adjacentInput
	}{
		{
			expected: []point{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}, {2, 2}},
			input: adjacentInput{
				start: point{1, 1},
				grid:  Grid{Matrix: [][]int32{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}},
			},
		},
		{
			expected: []point{{0, 0}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 2}},
			input: adjacentInput{
				start: point{1, 1},
				grid:  Grid{Matrix: [][]int32{{1, 0, 1}, {1, 1, 1}, {1, 0, 1}}},
			},
		},
		{
			expected: []point{{1, 1}, {0, 1}, {2, 2}},
			input: adjacentInput{
				start: point{1, 2},
				grid:  Grid{Matrix: [][]int32{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 1, 0}, {1, 0, 0, 0}}},
			},
		},
	}

	for _, tt := range tests {
		actual := adjacentFilledCell(tt.input.start, &tt.input.grid)
		assert.ElementsMatch(t, tt.expected, actual, tt.input.start)
	}

}

func TestBfs(t *testing.T) {
	type adjacentInput struct {
		grid  Grid
		start point
	}

	tests := []struct {
		expected int32
		input    adjacentInput
	}{
		{
			expected: 5,
			input: adjacentInput{
				start: point{0, 0},
				grid:  Grid{Matrix: [][]int32{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 1, 0}, {1, 0, 0, 0}}},
			},
		},
	}

	for _, tt := range tests {
		actual := bfs(tt.input.start, &tt.input.grid)
		assert.Equal(t, tt.expected, actual)
	}

}
