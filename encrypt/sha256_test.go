package encrypt

import "testing"

func TestSHA256Encrypt(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{ "t1", args{"aaaaaaaa"}, "1f3ce40415a2081fa3eee75fc39fff8e56c22270d1a978a7249b592dcebd20b4"},
		{ "t2", args{"bbbbbbbb"}, "fb398cc690e15ddba43ee811b6c0d3ec190901ad3df377fec9a1f9004b919a06"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA256Encrypt(tt.args.str); got != tt.want {
				t.Errorf("SHA256Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}