package cryptography

import (
	"testing"
)

func Test_generateRandomBytes(t *testing.T) {
	type args struct {
		bytes int
	}
	tests := []struct {
		name       string
		args       args
		wantLength int
		wantErr    bool
	}{
		{
			name: "Should be the right amount of bytes",
			args: args{
				bytes: 12,
			},
			wantLength: 12,
			wantErr:    false,
		},
		{
			name: "Should be the right amount of bytes",
			args: args{
				bytes: 8,
			},
			wantLength: 8,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateRandomBytes(tt.args.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateRandomBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantLength {
				t.Errorf("generateRandomBytes() got = %v, wantLength %v", got, tt.wantLength)
			}
		})
	}
}
