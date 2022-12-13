package cryptography

import (
	"bytes"
	"github.com/golang-jwt/jwt/v4"
	"reflect"
	"testing"
)

func TestCreateJWT(t *testing.T) {
	type args struct {
		key    []byte
		claims jwt.MapClaims
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Creates a jwt",
			args: args{
				key: []byte("testkey"),
				claims: jwt.MapClaims{
					"test": true,
				},
			},
			want:    "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ0ZXN0Ijp0cnVlfQ.YsdM6thJFAratz2a47d0gCd-EdoSg6HMIhV-R18JuRKS3j071Bc-qcqDcYc-aWaFJFBu06L87MvZksPKNTCRrQ",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateJWT(tt.args.key, tt.args.claims)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateJWT() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateJWT(t *testing.T) {
	type args struct {
		key         []byte
		signedToken string
	}
	tests := []struct {
		name    string
		args    args
		want    jwt.MapClaims
		wantErr bool
	}{
		{
			name: "Valid JWT",
			args: args{
				key:         []byte("testkey"),
				signedToken: "eyJhbGciOiJIUzUxMiJ9.eyJ0ZXN0Ijp0cnVlfQ.cBRaX4rlXOZb_40HKq4dUOnJOjRA6w0W6E0m7Pku5pLYdqwRBAZvYlnC_df5OGMI8eymRjiQJAMZz043qawNHg",
			},
			want: jwt.MapClaims{
				"test": true,
			},
			wantErr: false,
		},
		{
			name: "Invalid Signature",
			args: args{
				key:         []byte("testkey"),
				signedToken: "eyJhbGciOiJIUzUxMiJ9.eyJleHAiOjB9.Tk3TUQt6RI0cMjixkEOgSowmGjvPyHHWS7Gg-zbdvuzEwmflkYIFjCGMEY8Rj1EJZmXDUd7Bi2u_Ou9CcvEBqw",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid token",
			args: args{
				key:         []byte("testkey"),
				signedToken: "eyJhbGciOiJIUzUxMiJ9.eyJleHAiOjE2NzA5Njk1NTN9.AU04KX7-4Kxo7cENkmNKmG68zB0kZ2jvEKKrOX73xrjMfVTH1JtICFKIyZ00YVBTg9MEHC34VGk6dNNs-UT62g",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateJWT(tt.args.key, tt.args.signedToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateJWT() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getKeyFunc(t *testing.T) {
	type args struct {
		key []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Returns correct key",
			args: args{
				key: []byte("testkey"),
			},
			want: []byte("testkey"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getKeyFunc(tt.args.key)
			gotGot, err := got(nil)
			if err != nil {
				t.FailNow()
			}
			gotBytes, ok := gotGot.([]byte)
			if !ok {
				t.FailNow()
			}
			if bytes.Compare(gotBytes, tt.want) != 0 {
				t.Errorf("getKeyFunc()(nil).([]byte) = %v, want %v", gotBytes, tt.want)
			}
		})
	}
}
