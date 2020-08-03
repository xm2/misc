package threesum

import (
	"reflect"
	"testing"
)

func Test_threesum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][3]int
	}{
		// TODO: Add test cases.
		{"Test1", args{[]int{1, 0, -1}}, [][3]int{{-1, 0, 1}}},
		{"Test2", args{[]int{-4, -1, 0, 2, 1, 2}}, [][3]int{{-4, 2, 2}, {-1, 0, 1}}},
		{"Test3", args{[]int{-1, 0, 1, 2, -1, -4}}, [][3]int{{-1, -1, 2}, {-1, 0, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threesum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threesum() = %v, want %v", got, tt.want)
			}
		})
	}
}
