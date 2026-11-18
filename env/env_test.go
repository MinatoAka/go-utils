package env

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetEnv(t *testing.T) {
	// t.Setenv aman untuk testing, tidak bocor ke environment asli
	t.Setenv("TEST_KEY", "rahasia")

	tests := []struct {
		name    string
		key     string
		def     string
		want    string
	}{
		{"Ada di env", "TEST_KEY", "default", "rahasia"},
		{"Tidak ada di env", "MISSING_KEY", "default", "default"},
		{"Kosong tapi ada", "EMPTY_KEY", "default", "default"}, // Asumsi EMPTY_KEY tidak diset
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.key == "EMPTY_KEY" {
				t.Setenv("EMPTY_KEY", "")
			}
			got := GetEnv(tt.key, tt.def)
			if got != tt.want {
				t.Errorf("GetEnv() = %v, ingin %v", got, tt.want)
			}
		})
	}
}

func TestLoadDotEnv(t *testing.T) {
	// Bikin file .env sementara
	dir := t.TempDir()
	envPath := filepath.Join(dir, ".env")
	
	content := "DB_HOST=localhost\n# Ini komentar\nDB_PORT=5432\n"
	if err := os.WriteFile(envPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	if err := LoadDotEnv(envPath); err != nil {
		t.Fatalf("LoadDotEnv() error = %v", err)
	}

	if got := os.Getenv("DB_HOST"); got != "localhost" {
		t.Errorf("DB_HOST = %v, ingin localhost", got)
	}
	if got := os.Getenv("DB_PORT"); got != "5432" {
		t.Errorf("DB_PORT = %v, ingin 5432", got)
	}
}
