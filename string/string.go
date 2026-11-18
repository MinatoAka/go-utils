package string

import (
	"strconv"
	"strings"
)

// ToFloat64 mengkonversi string ke float64. Jika gagal, kembalikan defaultValue.
func ToFloat64(s string, defaultValue float64) float64 {
	val, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		return defaultValue
	}
	return val
}

// ToInt mengkonversi string ke int. Jika gagal, kembalikan defaultValue.
func ToInt(s string, defaultValue int) int {
	val, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		return defaultValue
	}
	return val
}

// Sanitize menghapus spasi di awal/akhir dan mengubah ke huruf kecil.
func Sanitize(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}
