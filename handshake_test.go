package hackerrankgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandshake(t *testing.T) {
	tests := []struct {
		name string
		n    int32
		want int32
	}{
		{
			name: "handshake",
			n:    1,
			want: 0,
		},
		{
			name: "handshake",
			n:    2,
			want: 1,
		},
		{
			name: "handshake",
			n:    3,
			want: 3,
		},
		{
			name: "handshake",
			n:    4,
			want: 6,
		},
	}
	for _, tt := range tests {
		actual := Handshake(tt.n)

		assert.Equal(t, tt.want, actual)
	}
}
