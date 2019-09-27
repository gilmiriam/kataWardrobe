package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWardrobeCombinator(t *testing.T) {
	tests := []struct {
		name string
		want [][]int
	}{
		{
			name: "A Combination using only the smallest element",
			want: [][]int{{50, 50, 50, 50, 50}},
		},

		{
			name: "A Combination using the two smallest elements",
			want: [][]int{{100, 100, 50}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WardrobeCombinator()
			if !assert.Contains(t, got, tt.want[0]) {
				t.Errorf("WardrobeCombinator() = %v, want %v", got, tt.want)
			}
		})
	}
}
