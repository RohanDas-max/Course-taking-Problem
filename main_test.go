package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoursesToTake(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		n     int
		want  []int
	}{
		{
			name:  "example 1",
			input: [][]int{{1, 0}},
			n:     2,
			want:  []int{0, 1},
		},
		{
			name:  "example 2",
			input: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			n:     4,
			want:  []int{0, 1, 2, 3},
		},
		{
			name:  "example 3",
			n:     1,
			input: [][]int{},
			want:  []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// got := coursesToTake(tt.input, tt.n) //brute Force problem
			got := findOrder(tt.input, tt.n) //Optimized solution using topological sorting
			assert.Equal(t, tt.want, got)
		})
	}
}
