package json

import (
	"path/filepath"
	"testing"
)

type dummyData struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}

func TestReadWriteFile(t *testing.T) {
	dir := t.TempDir()
	filePath := filepath.Join(dir, "test.json")

	original := dummyData{Name: "Budi", Age: 25, Email: "budi@mail.com"}

	// Test Write
	if err := WriteFile(filePath, original); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	// Test Read
	var result dummyData
	if err := ReadFile(filePath, &result); err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	if result.Name != original.Name || result.Age != original.Age || result.Email != original.Email {
		t.Errorf("ReadFile() = %v, ingin %v", result, original)
	}
}

func TestReadFile_NotExist(t *testing.T) {
	var result dummyData
	err := ReadFile("/path/that/does/not/exist.json", &result)
	if err == nil {
		t.Error("Ingin error, tapi dapat nil")
	}
}
