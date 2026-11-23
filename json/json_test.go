package json

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type dummyData struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}

func TestWriteAndRead(t *testing.T) {
	dir := t.TempDir()
	filePath := filepath.Join(dir, "test.json")

	original := dummyData{Name: "Budi", Age: 25, Email: "budi@mail.com"}

	// Test Write
	if err := Write(filePath, original); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	// Test Read
	result, err := Read[dummyData](filePath)
	if err != nil {
		t.Fatalf("Read() error = %v", err)
	}

	if result.Name != original.Name || result.Age != original.Age || result.Email != original.Email {
		t.Errorf("Read() = %v, ingin %v", result, original)
	}
}

func TestRead_NotExist(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "tidak-ada.json")
	_, err := Read[dummyData](filePath)
	if err == nil {
		t.Error("Ingin error, tapi dapat nil")
	}
}

func TestLoadOrCreate(t *testing.T) {
	dir := t.TempDir()
	filePath := filepath.Join(dir, "state.json")

	defaultData := dummyData{Name: "Default", Age: 10}

	// 1. File belum ada -> Harus buat baru dengan default
	result1, err := LoadOrCreate(filePath, defaultData)
	if err != nil {
		t.Fatalf("LoadOrCreate() error = %v", err)
	}
	if result1.Name != defaultData.Name || result1.Age != defaultData.Age {
		t.Errorf("LoadOrCreate() = %v, ingin %v", result1, defaultData)
	}

	// 2. File sudah ada -> Harus baca data yang ada (bukan default lagi)
	existingData := dummyData{Name: "Ubah", Age: 20}
	if err := Write(filePath, existingData); err != nil {
		t.Fatal(err)
	}

	result2, err := LoadOrCreate(filePath, defaultData)
	if err != nil {
		t.Fatalf("LoadOrCreate() error = %v", err)
	}
	if result2.Name != existingData.Name {
		t.Errorf("LoadOrCreate() = %v, ingin %v", result2, existingData)
	}
}

func TestUpdate(t *testing.T) {
	dir := t.TempDir()
	filePath := filepath.Join(dir, "update.json")

	original := dummyData{Name: "Budi", Age: 25}
	if err := Write(filePath, original); err != nil {
		t.Fatal(err)
	}

	// Update Age menjadi 30
	err := Update(filePath, func(d *dummyData) {
		d.Age = 30
	})
	if err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	// Verifikasi
	result, _ := Read[dummyData](filePath)
	if result.Age != 30 {
		t.Errorf("Setelah Update(), Age = %d, ingin 30", result.Age)
	}
	if result.Name != "Budi" {
		t.Errorf("Setelah Update(), Name berubah jadi %s", result.Name)
	}
}

func TestPrettyFunctions(t *testing.T) {
	dir := t.TempDir()
	filePath := filepath.Join(dir, "pretty.json")

	original := dummyData{Name: "Budi", Age: 25, Email: "budi@mail.com"}

	// 1. Test MarshalPretty (harus menghasilkan string dengan newline)
	str, err := MarshalPretty(original)
	if err != nil {
		t.Fatalf("MarshalPretty() error = %v", err)
	}
	if !strings.Contains(str, "\n") {
		t.Error("MarshalPretty() tidak menghasilkan format rapi (tidak ada newline)")
	}

	// 2. Test WritePretty (simpan ke file)
	if err := WritePretty(filePath, original); err != nil {
		t.Fatalf("WritePretty() error = %v", err)
	}

	// Baca raw file untuk memastikan ada indentasi
	rawData, _ := os.ReadFile(filePath)
	if !strings.Contains(string(rawData), "\n") {
		t.Error("WritePretty() menulis file tanpa format rapi")
	}
}