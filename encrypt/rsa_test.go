package encrypt

import (
	"crypto/rsa"
	"testing"
)

var (
	pubicKey   = []byte("-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA+g025mP2iUqgXOiGK5pJ\nh1PGRYuR6kzyHZpvfOls6Fbga+xP48fgrCzHIatCj8yu0I2yilXaGUJ3MHMNElzt\nJp02NEW/NKQNR6q3R7j21V98cAQaJ/Hgauf0ehSop2KKPHXb6vBasmGNZgc3iGEg\nL0p2amlnFvUmcvv/EQuabASlBrMzSAQx+85LCci9LM6SzCxbwCQAUHARfg78zop2\nbPdDIV3qeTQXrcI6MtGidzmLCxWpAShpliEugPsRRCrQV97vMEn4CgJPFnnBQ8Cp\n5QVMUciYWAkBv0WqqPI0dPCBjYasVJqwKfMvY+xi7SdJk6BjM4j44auUKpjh4jax\nVQIDAQAB\n-----END PUBLIC KEY-----\n")
	privateKey = []byte("-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQD6DTbmY/aJSqBc\n6IYrmkmHU8ZFi5HqTPIdmm986WzoVuBr7E/jx+CsLMchq0KPzK7QjbKKVdoZQncw\ncw0SXO0mnTY0Rb80pA1HqrdHuPbVX3xwBBon8eBq5/R6FKinYoo8ddvq8FqyYY1m\nBzeIYSAvSnZqaWcW9SZy+/8RC5psBKUGszNIBDH7zksJyL0szpLMLFvAJABQcBF+\nDvzOinZs90MhXep5NBetwjoy0aJ3OYsLFakBKGmWIS6A+xFEKtBX3u8wSfgKAk8W\necFDwKnlBUxRyJhYCQG/Raqo8jR08IGNhqxUmrAp8y9j7GLtJ0mToGMziPjhq5Qq\nmOHiNrFVAgMBAAECggEALXjluPQIy2LE0uB+NcOzB2cWsWL2QBOKHBQPnjmc0oup\n+DZbOOMKWoUyGU9ZpWjliYubIkkGhn1ty2PAxswbXGBOkSb15I7I03WaKvLaAGeC\n66PzVbxIwWg5L1I+cWIcHW6ZdZoMBFB1NE6vrJ9tCOyTqDpOwwAF4/crpV/kzyuV\nCX9lAHEKzt5XOEu2zXWPtl5q6bku3h7WZqqDeGvqRzxn76KNBqihKbQLyy7I3a8N\nqbqy2JY2GRBuMpmDKmdxSwFw8ZdPpF3IlbJtyzpQee4g64GVugjleQz7gPKMFdV0\nVNHoP95qwGwHjGrt6KdddBQF3zKPlH8g66jO5H1Z6QKBgQD+Fz4D2l8CLDMina+U\nFu+WQqpEmNJtzpsC/yWtykWJzmgUV88xkBTFzfh0FO8TCBMhNMMZd1tIllSvMTR6\nBhffNsq+pJLIfdUP2KXLeOKU/gkfisbsUFnTUk5g/3fTD+mSBw2o+G7zfEmJpQbo\nj+M9SPySVqfscnS3I14XPkDdrwKBgQD77jPf/5SIyJQyShjAEMA+p4jZwqknzxTB\nZI6d3x8wr8WLlJ67dlcLBo4n6SkOlh3g3Pk76Pp67jd78EgI9cJE5Fw8KnK5PfVg\n2PqBoqeHTNStG3+a2qYKmROLLYlk88A4NUXTE9q+YPY/BKdOucZYFdpgXeEeJ3+b\nijihtuaGOwKBgQCT5e14CqzQs++Tz5s6pNsSaH15bkbWGbOK2/8PmLQ/UYCtjqsQ\nm2Ar/wcGcoTyW/nekzqY4SumJbLnQsnb2R4eFGOQ8nRp9SyGi8F9nUCuHuwqivnc\n5axTYA41UZ5qoVGAaVCfMPRMD+UyFHff+Jj5vDkf6QbXj4u2agTDnL3fCwKBgGZW\nZuAxNCc0Q28CRneO1mkdKSw8NPak2iuNK5nZzxO58yc2IhzIf6dM7GYuAq2YMV7s\npJm/lDBnoEZKXxi0rEvL7+PI4n15O7oxqeELL754aVRAfV9sPlLDx/qbbtqBYa3z\nhf+uawc24BDNVSFXuciCaAkJWa8kIGQQi9y1LYONAoGAF+0ckKWMP6L/PjTNJqH4\nGcre44C33gpBgEbTTyIgMCQ/c+yGZXPIbjKD7r2X2YT3+XxR3gFnDLxU+kkByIIB\nxAHLDL7DT8heeG/JSTCZJcQoyeA06ws+Bbu60QhRZaHEbCNBocM+xX90taovzhX7\ncDpMtXYEWpbMWlNeR49vfiY=\n-----END PRIVATE KEY-----\n")
)

func TestNewXRsa(t *testing.T) {
	type args struct {
		publicKey  []byte
		privateKey []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"t1",
			args{
				publicKey:  pubicKey,
				privateKey: privateKey,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewXRsa(tt.args.publicKey, tt.args.privateKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewXRsa() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_xrsa_PublicEncrypt(t *testing.T) {
	type fields struct {
		publicKey  *rsa.PublicKey
		privateKey *rsa.PrivateKey
	}
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &xrsa{
				publicKey:  tt.fields.publicKey,
				privateKey: tt.fields.privateKey,
			}
			got, err := r.PublicEncrypt(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("PublicEncrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PublicEncrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}