package addtwonumbers

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"sum < 10", args{makeNodeListFromNum(1), makeNodeListFromNum(2)}, makeNodeListFromNum(3)},
		{"sum = 10", args{makeNodeListFromNum(8), makeNodeListFromNum(2)}, makeNodeListFromNum(10)},
		{"more than 3-digtal", args{makeNodeListFromNum(3), makeNodeListFromNum(212)}, makeNodeListFromNum(3 + 212)},
		{"big number", args{makeNodeListFromNum(334499), makeNodeListFromNum(9876512)}, makeNodeListFromNum(334499 + 9876512)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
			if got := addTwoNumbersV2(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
