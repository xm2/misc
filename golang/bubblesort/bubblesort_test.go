package bubblesort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubblesort(t *testing.T) {
	type args struct {
		parr *[]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"bsort1", args{&[]int{1, 2, 3, 4, 5}}, []int{1, 2, 3, 4, 5}},
		{"bsort2", args{&[]int{5, 4, 3, 2, 1}}, []int{1, 2, 3, 4, 5}},
		{"bsort3", args{&[]int{5, 1, 3, 2, 4}}, []int{1, 2, 3, 4, 5}},
		{"bsort4", args{&[]int{5, 1}}, []int{1, 5}},
		{"bsort5", args{&[]int{5}}, []int{5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Bubblesort(tt.args.parr)
			assert.Equal(t, tt.want, *tt.args.parr)
		})
	}
}
