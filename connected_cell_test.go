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
		{
			name:  "ConnectedCell",
			input: [][]int32{{0, 0}},
			want:  0,
		},
		{
			name:  "ConnectedCell",
			input: [][]int32{{0, 0}, {0, 1}},
			want:  1,
		},
		{
			name:  "ConnectedCell",
			input: [][]int32{{1, 1}, {1, 1}},
			want:  4,
		},
	}

	for _, tt := range tests {
		actual := connectedCell(tt.input)

		assert.Equal(t, tt.want, actual)
	}
}

func TestAdjacentFilledCell(t *testing.T) {
	type adjacentInput struct {
		grid  Grid
		start Cell
	}

	tests := []struct {
		expected []Cell
		input    adjacentInput
	}{
		{
			expected: []Cell{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}, {2, 2}},
			input: adjacentInput{
				start: Cell{1, 1},
				grid:  [][]int32{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			},
		},
		{
			expected: []Cell{{0, 0}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 2}},
			input: adjacentInput{
				start: Cell{1, 1},
				grid:  [][]int32{{1, 0, 1}, {1, 1, 1}, {1, 0, 1}},
			},
		},
		{
			expected: []Cell{{1, 1}, {0, 1}, {2, 2}},
			input: adjacentInput{
				start: Cell{y: 1, x: 2},
				grid:  [][]int32{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 1, 0}, {1, 0, 0, 0}},
			},
		},
	}

	for _, tt := range tests {
		actual := adjacentFilledCell(tt.input.start, tt.input.grid)
		assert.ElementsMatch(t, tt.expected, actual, tt.input.start)
	}

}

func TestBfs(t *testing.T) {
	type adjacentInput struct {
		grid  Grid
		start Cell
	}

	tests := []struct {
		expected []Cell
		input    adjacentInput
	}{
		{
			expected: []Cell{{0, 0}, {0, 1}, {1, 1}, {1, 2}, {2, 2}},
			input: adjacentInput{
				start: Cell{0, 0},
				grid: [][]int32{
					{1, 1, 0, 0},
					{0, 1, 1, 0},
					{0, 0, 1, 0},
					{1, 0, 0, 0},
				},
			},
		},
	}

	for _, tt := range tests {
		actual := bfs(tt.input.start, tt.input.grid)
		assert.ElementsMatch(t, tt.expected, actual)
	}

}

func TestPop(t *testing.T) {
	queue := []Cell{
		{1, 0},
		{2, 2},
		{5, 2},
	}

	result, q := pop(queue)

	assert.Equal(t, Cell{1, 0}, result)
	assert.ElementsMatch(t, []Cell{{2, 2}, {5, 2}}, q)
}
