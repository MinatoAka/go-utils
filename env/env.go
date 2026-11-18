package env

import (
	"bufio"
	"os"
	"strings"
)

// GetEnv mengambil value dari environment variable. 
// Jika kosong, kembalikan defaultValue.
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// GetEnvInt sama seperti GetEnv tapi untuk integer.
func GetEnvInt(key string, defaultValue int) int {
	// (Bisa ditambahin logic strconv.Atoi di sini)
	return defaultValue 
}

// LoadDotEnv sederhana untuk baca file .env di root project (tanpa dependency eksternal).
func LoadDotEnv(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			os.Setenv(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
		}
	}
	return scanner.Err()
}
