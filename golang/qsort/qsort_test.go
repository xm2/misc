package qsort

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func generateSlice(n int) []int {
	rand.Seed(time.Now().UnixNano())
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = rand.Intn(65535)
	}

	return s
}

func TestQsort(t *testing.T) {
	type args struct {
		parr *[]int
		min  int
		max  int
	}
	stubData1 := generateSlice(1)
	stubData2 := generateSlice(2)
	stubData3 := generateSlice(10)
	stubData4 := generateSlice(100)
	stubData5 := generateSlice(1000)
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"qsort1", args{&stubData1, 0, len(stubData1) - 1}},
		{"qsort2", args{&stubData2, 0, len(stubData2) - 1}},
		{"qsort3", args{&stubData3, 0, len(stubData3) - 1}},
		{"qsort4", args{&stubData4, 0, len(stubData4) - 1}},
		{"qsort5", args{&stubData5, 0, len(stubData5) - 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Qsort(tt.args.parr, tt.args.min, tt.args.max)
			assert.Equal(t, true, sort.IntsAreSorted(*tt.args.parr))
		})
	}
}

func TestQsort2(t *testing.T) {
	type args struct {
		parr *[]int
		min  int
		max  int
	}
	stubData1 := generateSlice(1)
	stubData2 := generateSlice(2)
	stubData3 := generateSlice(10)
	stubData4 := generateSlice(100)
	stubData5 := generateSlice(1000)
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"qsort1", args{&stubData1, 0, len(stubData1) - 1}},
		{"qsort2", args{&stubData2, 0, len(stubData2) - 1}},
		{"qsort3", args{&stubData3, 0, len(stubData3) - 1}},
		{"qsort4", args{&stubData4, 0, len(stubData4) - 1}},
		{"qsort5", args{&stubData5, 0, len(stubData5) - 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Qsort2(tt.args.parr, tt.args.min, tt.args.max)
			assert.Equal(t, true, sort.IntsAreSorted(*tt.args.parr))
		})
	}
}

func BenchmarkQsort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		stubData := generateSlice(10000)
		b.StartTimer()
		Qsort(&stubData, 0, len(stubData)-1)
	}
}
func BenchmarkQsort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		stubData := generateSlice(10000)
		b.StartTimer()
		Qsort2(&stubData, 0, len(stubData)-1)
	}
}
