package timex

import (
	"time"
)

const (
	// Layout standar untuk database/log
	LayoutDateTime = "2006-01-02 15:04:05"
	// Layout format tanggal umum di Indonesia
	LayoutDateID = "02/01/2006 15:04"
	// Layout hanya jam
	LayoutTimeOnly = "15:04:05"
)

// FormatDateTime memformat waktu ke format standar "2006-01-02 15:04:05".
func FormatDateTime(t time.Time) string {
	return t.Format(LayoutDateTime)
}

// FormatDateID memformat waktu ke format umum Indonesia "02/01/2006 15:04".
func FormatDateID(t time.Time) string {
	return t.Format(LayoutDateID)
}

// FormatTimeOnly memformat waktu hanya jam, menit, detik "15:04:05".
func FormatTimeOnly(t time.Time) string {
	return t.Format(LayoutTimeOnly)
}

// NowDateTime adalah shortcut untuk FormatDateTime(time.Now()).
func NowDateTime() string {
	return FormatDateTime(time.Now())
}
