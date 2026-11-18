package console

import (
	"testing"
	"time"
)

func TestFmtTypeWriter(t *testing.T) {
	// Gunakan speed sangat cepat agar test selesai instan
	// dan pastikan tidak terjadi panic
	FmtTypeWriter("Testing typewriter effect!", 5*time.Nanosecond)
	FmtTypeWriter("Tanpa newline di akhir", 5*time.Nanosecond)
	FmtTypeWriter("Dengan newline di akhir\n", 5*time.Nanosecond)
}

func TestFmtLogWriter(t *testing.T) {
	// Pastikan fungsi berjalan tanpa error/panic
	FmtLogWriter("Testing log writer dengan timestamp!", 5*time.Nanosecond)
}
