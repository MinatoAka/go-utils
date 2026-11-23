package json

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Read membaca file JSON dan mem-parsing-nya langsung dari stream (hemat RAM).
func Read[T any](path string) (T, error) {
	var result T
	file, err := os.Open(path)
	if err != nil {
		return result, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&result)
	return result, err
}

// Write menulis struct ke file JSON dengan kecepatan maksimal (Atomic Write).
// File JSON akan berupa satu baris panjang (tanpa pretty-print) agar sangat cepat.
func Write[T any](path string, v T) error {
	data, err := json.Marshal(v) // <-- Pakai Marshal, bukan MarshalIndent
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	// Gunakan permission 0700 agar hanya user yang bisa akses foldernya
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}

	tmp, err := os.CreateTemp(dir, ".json-*.tmp")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	defer os.Remove(tmpName)

	// Set permission file ke 0600 (rahasia)
	if err := tmp.Chmod(0600); err != nil {
		tmp.Close()
		return err
	}

	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return err
	}
	
	// HAPUS tmp.Sync() DEMI KECEPATAN MAKSIMAL
	tmp.Close()

	return os.Rename(tmpName, path)
}

// LoadOrCreate membaca file, atau membuat file baru dengan nilai default jika belum ada.
// Sangat berguna untuk inisialisasi state/config di awal bot.
func LoadOrCreate[T any](path string, defaultVal T) (T, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// File belum ada, tulis nilai default
		if err := Write(path, defaultVal); err != nil {
			return defaultVal, err
		}
		return defaultVal, nil
	}
	// File sudah ada, baca
	return Read[T](path)
}

// Update adalah fungsi "Read-Modify-Write" untuk mengubah sebagian data dengan aman.
func Update[T any](path string, updateFn func(*T)) error {
	// 1. Baca data saat ini
	data, err := Read[T](path)
	if err != nil {
		return err
	}

	// 2. Modifikasi data via callback
	updateFn(&data)

	// 3. Simpan kembali (Atomic)
	return Write(path, data)
}

// MarshalPretty mengubah struct menjadi string JSON yang rapi (indented).
// Berguna untuk keperluan print ke terminal atau log.
func MarshalPretty[T any](v T) (string, error) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// WritePretty menulis struct ke file JSON dengan format rapi (indented).
// Gunakan ini hanya untuk file yang butuh dibaca manusia (misal config atau debug).
func WritePretty[T any](path string, v T) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}

	tmp, err := os.CreateTemp(dir, ".json-*.tmp")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	defer os.Remove(tmpName)

	if err := tmp.Chmod(0600); err != nil {
		tmp.Close()
		return err
	}

	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return err
	}
	
	tmp.Close()
	return os.Rename(tmpName, path)
}