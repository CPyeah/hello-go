package algorithm

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		sorted []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 3, 5, 6, 7}, 4}, -1},
		{"test2", args{[]int{1, 2, 3, 5, 6, 7}, 3}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.sorted, tt.args.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
