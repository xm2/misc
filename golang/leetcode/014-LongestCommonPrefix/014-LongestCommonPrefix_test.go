package longestcommonprefix

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Test1", args{strs: []string{"flower", "flow", "flight"}}, "fl"},
		{"Test2", args{strs: []string{"dog", "racecar", "car"}}, ""},
		{"Test3", args{strs: []string{"hotdog", "hotwind", "hospital"}}, "ho"},
		{"Test4", args{strs: []string{"hot", "hotwind", "hotwater", "hotdog"}}, "hot"},
		{"Test5", args{strs: []string{"hotdog", "hotwind", "hotwater", "hot"}}, "hot"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
