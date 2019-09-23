package main

import (
	"reflect"
	"testing"
)

func TestWardrobeCombinator(t *testing.T) {
	tests := []struct {
		name string
		want combinator
	}{
		{
			name: "A Combination using only the smallest element",
			want: combinator{
				[]int{50, 50, 50, 50, 50},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WardrobeCombinator()
			if reflect.DeepEqual(got, tt.want) != true {
				t.Errorf("WardrobeCombinator() = %v, want %v", got, tt.want)
			}
		})
	}
}
