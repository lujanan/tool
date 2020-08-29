package util

import (
	"fmt"
	"testing"
)

func TestRandString(t *testing.T) {
	type args struct {
		n      int
		source string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t0", args{4, ""}, 4},
		{"t1", args{6, ""}, 6},
		{"t2", args{8, ""}, 8},
		{"t3", args{10, ""}, 10},
		{"t4", args{20, ""}, 20},
		{"t5", args{30, ""}, 30},
		{"t6", args{40, ""}, 40},
		{"t7", args{50, ""}, 50},
		{"t8", args{60, ""}, 60},
		{"t9", args{70, ""}, 70},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandString(tt.args.n, tt.args.source); len(got) != tt.want {
				t.Errorf("RandString() = %v, want %v", got, tt.want)
			} else {
				fmt.Printf("%s, want string lenght: [%d], result: [%s], lenght: [%d]", tt.name, tt.want, got, len(got))
			}
		})
	}
}
