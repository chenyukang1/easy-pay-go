package xrsa

import (
	"crypto/rsa"
	"testing"
)

func TestSign(t *testing.T) {
	type args struct {
		bytes      []byte
		privateKey *rsa.PrivateKey
		signType   SignType
	}
	tests := []struct {
		name          string
		args          args
		wantSignature string
		wantErr       bool
	}{
		{
			name: "test nil",
			args: args{
				bytes:      nil,
				privateKey: nil,
				signType:   RSA2,
			},
			wantSignature: "",
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSignature, err := Sign(tt.args.bytes, tt.args.privateKey, tt.args.signType)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSignature != tt.wantSignature {
				t.Errorf("Sign() gotSignature = %v, want %v", gotSignature, tt.wantSignature)
			}
		})
	}
}
