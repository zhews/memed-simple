package meme

import (
	"os"
	"reflect"
	"testing"
)

func TestParseFromEnvironmentalVariables(t *testing.T) {
	tests := []struct {
		name      string
		env       map[string]string
		createDir bool
		dir       string
		want      Config
		wantErr   bool
	}{
		{
			name: "Proper Configuration",
			env: map[string]string{
				"MEMED_PORT":               "8080",
				"MEMED_DATABASE_URL":       "postgresql://localhost:5432/memed",
				"MEMED_CORS_ALLOW_ORIGINS": "http://localhost:4200",
				"MEMED_MEME_DIRECTORY":     "tmp",
				"MEMED_USER_MICROSERVICE":  "http://user:7070",
				"MEMED_USER_ENDPOINT":      "/user",
				"MEMED_ACCESS_SECRET_KEY":  "access_secret_key",
			},
			createDir: true,
			dir:       "tmp",
			want: Config{
				Port:             8080,
				DatabaseURL:      "postgresql://localhost:5432/memed",
				CorsAllowOrigins: "http://localhost:4200",
				UserMicroservice: "http://user:7070",
				UserEndpoint:     "/user",
				MemeDirectory:    "tmp",
				AccessSecretKey:  "access_secret_key",
			},
			wantErr: false,
		},
		{
			name: "Invalid Port",
			env: map[string]string{
				"MEMED_PORT": "abcd",
			},
			want:    Config{},
			wantErr: true,
		},
		{
			name: "Invalid Meme Directory",
			env: map[string]string{
				"MEMED_PORT":           "8080",
				"MEMED_MEME_DIRECTORY": "tmp",
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
		if tt.createDir {
			err := os.Mkdir(tt.dir, 0755)
			if err != nil {
				t.FailNow()
			}
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseFromEnvironmentalVariables()
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFromEnvironmentalVariables() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFromEnvironmentalVariables() got = %v, want %v", got, tt.want)
			}
		})
		for key := range tt.env {
			err := os.Unsetenv(key)
			if err != nil {
				t.FailNow()
			}
		}
		if tt.createDir {
			err := os.Remove(tt.dir)
			if err != nil {
				t.FailNow()
			}
		}
	}
}
