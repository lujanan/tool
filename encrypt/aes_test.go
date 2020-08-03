package encrypt

import (
	"encoding/base64"
	"reflect"
	"testing"
)

func TestAesCBCEncrypt(t *testing.T) {
	type args struct {
		rawData []byte
		key     []byte
		iv      []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"t1",
			args{
				rawData: []byte("hello world"),
				key:     []byte("aaaaaaaaaaaaaaaa"),
				iv:      []byte("1111111111111111"),
			},
			"tBSp2FoyQiI+GWJS5x+z0w==",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AesCBCEncrypt(tt.args.rawData, tt.args.key, tt.args.iv)
			if (err != nil) != tt.wantErr {
				t.Errorf("AesCBCEncrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			str := base64.StdEncoding.EncodeToString(got)
			if !reflect.DeepEqual(str, tt.want) {
				t.Errorf("AesCBCEncrypt() got = %v (string = %v), want %v", got, str, tt.want)
			}
		})
	}
}

func TestAesCBCDecrypt(t *testing.T) {
	type args struct {
		encryptData string
		key         []byte
		iv          []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"t1",
			args{
				encryptData: "tBSp2FoyQiI+GWJS5x+z0w==",
				key:         []byte("aaaaaaaaaaaaaaaa"),
				iv:          []byte("1111111111111111"),
			},
			"hello world",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := base64.StdEncoding.DecodeString(tt.args.encryptData)
			if err != nil {
				t.Errorf("AesCBCDecrypt() error = %v, encryptData %v", err, tt.args.encryptData)
				return
			}
			got, err := AesCBCDecrypt(data, tt.args.key, tt.args.iv)
			if (err != nil) != tt.wantErr {
				t.Errorf("AesCBCDecrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(string(got), tt.want) {
				t.Errorf("AesCBCDecrypt() got = %v (string = %v), want %v", got, string(got), tt.want)
			}
		})
	}
}
