package medianoftwosortedarrays

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"Test1", args{[]int{1, 9}, []int{10}}, 9.0},
		{"Test2", args{[]int{1}, []int{9, 10}}, 9.0},
		{"Test3", args{[]int{1}, []int{10}}, 5.5},
		{"Test4", args{[]int{1}, nil}, 1},
		{"Test5", args{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}}, 5.5},
		{"Test6", args{[]int{1, 3, 5, 7, 9}, []int{2, 4, 6, 8, 10}}, 5.5},
		{"Test7", args{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9}}, 5},
		{"Test8", args{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8, 9}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
