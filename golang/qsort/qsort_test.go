package qsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQsort(t *testing.T) {
	type args struct {
		parr *[]int
		min  int
		max  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"qsort1", args{&[]int{1, 2, 3, 4, 5}, 0, 4}, []int{1, 2, 3, 4, 5}},
		{"qsort2", args{&[]int{5, 4, 3, 2, 1}, 0, 4}, []int{1, 2, 3, 4, 5}},
		{"qsort3", args{&[]int{5, 1, 3, 2, 4}, 0, 4}, []int{1, 2, 3, 4, 5}},
		{"qsort4", args{&[]int{5, 1}, 0, 1}, []int{1, 5}},
		{"qsort5", args{&[]int{5}, 0, 0}, []int{5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Qsort(tt.args.parr, tt.args.min, tt.args.max)
			assert.Equal(t, tt.want, *tt.args.parr)
		})
	}
}

func TestQsort2(t *testing.T) {
	type args struct {
		parr *[]int
		min  int
		max  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"qsort1", args{&[]int{1, 2, 3, 4, 5}, 0, 4}, []int{1, 2, 3, 4, 5}},
		{"qsort2", args{&[]int{5, 4, 3, 2, 1}, 0, 4}, []int{1, 2, 3, 4, 5}},
		{"qsort3", args{&[]int{5, 1, 3, 2, 4}, 0, 4}, []int{1, 2, 3, 4, 5}},
		{"qsort4", args{&[]int{5, 1}, 0, 1}, []int{1, 5}},
		{"qsort5", args{&[]int{5}, 0, 0}, []int{5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Qsort2(tt.args.parr, tt.args.min, tt.args.max)
			assert.Equal(t, tt.want, *tt.args.parr)
		})
	}
}
