package main

import (
	"fmt"
	"reflect"
	"testing"
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
			includes := 0
			for _, comb := range got {
				fmt.Println(comb, " => ", tt.want[0])
				if reflect.DeepEqual(comb, tt.want[0]) {
					includes = 1
					fmt.Println(includes)
				}
			}

			if includes == 0 {
				t.Errorf("WardrobeCombinator() = %v, want %v", got, tt.want)
			}
		})
	}
}
