package timex

import (
	"testing"
	"time"
)

func TestFormatDateTime(t *testing.T) {
	// Pakai waktu fix biar tes nggak gagal kalau dijalankan di waktu berbeda
	fixedTime := time.Date(2026, 11, 19, 10, 30, 45, 0, time.UTC)
	
	want := "2026-11-19 10:30:45"
	got := FormatDateTime(fixedTime)
	
	if got != want {
		t.Errorf("FormatDateTime() = %q, ingin %q", got, want)
	}
}

func TestFormatDateID(t *testing.T) {
	fixedTime := time.Date(2026, 11, 19, 10, 30, 0, 0, time.UTC)
	
	want := "19/11/2026 10:30"
	got := FormatDateID(fixedTime)
	
	if got != want {
		t.Errorf("FormatDateID() = %q, ingin %q", got, want)
	}
}

func TestFormatTimeOnly(t *testing.T) {
	fixedTime := time.Date(2026, 11, 19, 15, 4, 5, 0, time.UTC)
	
	want := "15:04:05"
	got := FormatTimeOnly(fixedTime)
	
	if got != want {
		t.Errorf("FormatTimeOnly() = %q, ingin %q", got, want)
	}
}
