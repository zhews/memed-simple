package user

import (
	"github.com/zhews/memed-simple/pkg/cryptography"
	"os"
	"reflect"
	"testing"
)

func TestFromEnvironmentalVariables(t *testing.T) {
	tests := []struct {
		name    string
		env     map[string]string
		want    Config
		wantErr bool
	}{
		{
			name: "Proper Configuration",
			env: map[string]string{
				"MEMED_PORT":                          "7070",
				"MEMED_CORS_ALLOW_ORIGINS":            "http://localhost:4200",
				"MEMED_DATABASE_URL":                  "postgresql://localhost:5432/memed",
				"MEMED_BASE_URI":                      "api.memed.io",
				"MEMED_ACCESS_SECRET_KEY":             "access_secret_key",
				"MEMED_ACCESS_TOKEN_VALID_SECONDS":    "10",
				"MEMED_REFRESH_SECRET_KEY":            "refresh_secret_key",
				"MEMED_REFRESH_TOKEN_VALID_HOURS":     "24",
				"MEMED_ENCRYPTION_KEY":                "encryption_key",
				"MEMED_ARGON2ID_PARAMETER_SALT_SIZE":  "16",
				"MEMED_ARGON2ID_PARAMETER_ITERATIONS": "2",
				"MEMED_ARGON2ID_PARAMETER_MEMORY":     "1024",
				"MEMED_ARGON2ID_PARAMETER_THREADS":    "2",
				"MEMED_ARGON2ID_PARAMETER_KEY_LENGTH": "32",
			},
			want: Config{
				Port:                    7070,
				CorsAllowOrigins:        "http://localhost:4200",
				DatabaseURL:             "postgresql://localhost:5432/memed",
				BaseURI:                 "api.memed.io",
				AccessSecretKey:         "access_secret_key",
				AccessTokenValidSeconds: 10,
				RefreshSecretKey:        "refresh_secret_key",
				RefreshTokenValidHours:  24,
				EncryptionKey:           "encryption_key",
				Argon2IDParameter: cryptography.Argon2IDParameter{
					SaltSize:   16,
					Iterations: 2,
					Memory:     1024,
					Threads:    2,
					KeyLength:  32,
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid port",
			env: map[string]string{
				"MEMED_PORT": "abcd",
			},
			want:    Config{},
			wantErr: true,
		},
		{
			name: "Invalid access token valid seconds",
			env: map[string]string{
				"MEMED_PORT":                       "8080",
				"MEMED_ACCESS_TOKEN_VALID_SECONDS": "abcd",
			},
			want:    Config{},
			wantErr: true,
		},
		{
			name: "Invalid refresh token valid hours",
			env: map[string]string{
				"MEMED_PORT":                       "8080",
				"MEMED_ACCESS_TOKEN_VALID_SECONDS": "10",
				"MEMED_REFRESH_TOKEN_VALID_HOURS":  "abcd",
			},
			want:    Config{},
			wantErr: true,
		},
		{
			name: "Invalid argon2id parameter salt size",
			env: map[string]string{
				"MEMED_PORT":                         "8080",
				"MEMED_ACCESS_TOKEN_VALID_SECONDS":   "10",
				"MEMED_REFRESH_TOKEN_VALID_HOURS":    "24",
				"MEMED_ARGON2ID_PARAMETER_SALT_SIZE": "abcd",
			},
			want:    Config{},
			wantErr: true,
		},
		{
			name: "Invalid argon2id parameter iterations",
			env: map[string]string{
				"MEMED_PORT":                          "8080",
				"MEMED_ACCESS_TOKEN_VALID_SECONDS":    "10",
				"MEMED_REFRESH_TOKEN_VALID_HOURS":     "24",
				"MEMED_ARGON2ID_PARAMETER_SALT_SIZE":  "16",
				"MEMED_ARGON2ID_PARAMETER_ITERATIONS": "abcd",
			},
			want:    Config{},
			wantErr: true,
		},
		{
			name: "Invalid argon2id parameter memory",
			env: map[string]string{
				"MEMED_PORT":                          "8080",
				"MEMED_ACCESS_TOKEN_VALID_SECONDS":    "10",
				"MEMED_REFRESH_TOKEN_VALID_HOURS":     "24",
				"MEMED_ARGON2ID_PARAMETER_SALT_SIZE":  "16",
				"MEMED_ARGON2ID_PARAMETER_ITERATIONS": "2",
				"MEMED_ARGON2ID_PARAMETER_MEMORY":     "abcd",
			},
			want:    Config{},
			wantErr: true,
		},
		{
			name: "Invalid argon2id parameter threads",
			env: map[string]string{
				"MEMED_PORT":                          "8080",
				"MEMED_ACCESS_TOKEN_VALID_SECONDS":    "10",
				"MEMED_REFRESH_TOKEN_VALID_HOURS":     "24",
				"MEMED_ARGON2ID_PARAMETER_SALT_SIZE":  "16",
				"MEMED_ARGON2ID_PARAMETER_ITERATIONS": "2",
				"MEMED_ARGON2ID_PARAMETER_MEMORY":     "1024",
				"MEMED_ARGON2ID_PARAMETER_THREADS":    "abc",
			},
			want:    Config{},
			wantErr: true,
		},
		{
			name: "Invalid argon2id parameter key length",
			env: map[string]string{
				"MEMED_PORT":                          "8080",
				"MEMED_ACCESS_TOKEN_VALID_SECONDS":    "10",
				"MEMED_REFRESH_TOKEN_VALID_HOURS":     "24",
				"MEMED_ARGON2ID_PARAMETER_SALT_SIZE":  "16",
				"MEMED_ARGON2ID_PARAMETER_ITERATIONS": "2",
				"MEMED_ARGON2ID_PARAMETER_MEMORY":     "1024",
				"MEMED_ARGON2ID_PARAMETER_THREADS":    "2",
				"MEMED_ARGON2ID_PARAMETER_KEY_LENGTH": "abcd",
			},
			want:    Config{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		for key, value := range tt.env {
			err := os.Setenv(key, value)
			if err != nil {
				t.FailNow()
			}
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromEnvironmentalVariables()
			if (err != nil) != tt.wantErr {
				t.Errorf("FromEnvironmentalVariables() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromEnvironmentalVariables() got = %v, want %v", got, tt.want)
			}
		})
		for key := range tt.env {
			err := os.Unsetenv(key)
			if err != nil {
				t.FailNow()
			}
		}
	}
}
