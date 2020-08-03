package encrypt

import "testing"

func TestMd5Encrypt(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{ "t1", args{"aaaaaaaa"}, "3dbe00a167653a1aaee01d93e77e730e"},
		{ "t2", args{"bbbbbbbb"}, "810247419084c82d03809fc886fedaad"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5Encrypt(tt.args.str); got != tt.want {
				t.Errorf("Md5Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}