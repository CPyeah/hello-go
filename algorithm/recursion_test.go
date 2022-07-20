package algorithm

import "testing"

func Test_fibonacci(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test-1", args{1}, 1},
		{"test-2", args{2}, 1},
		{"test-3", args{3}, 2},
		{"test-4", args{4}, 3},
		{"test-5", args{5}, 5},
		{"test-6", args{6}, 8},
		{"test-7", args{7}, 13},
		{"test-45", args{45}, 1134903170},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacci(tt.args.num); got != tt.want {
				t.Errorf("fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
