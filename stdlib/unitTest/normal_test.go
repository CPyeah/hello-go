package unitTest

import "testing"

func Test_isNegative(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"testName0", args{0}, false},
		{"testName1", args{2}, false},
		{"testName2", args{-1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNegative(tt.args.num); got != tt.want {
				t.Errorf("isNegative() = %v, want %v", got, tt.want)
			}
		})
	}
}
