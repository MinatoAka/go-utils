package currency

import "testing"

func TestFormatRupiah(t *testing.T) {
	tests := []struct {
		input int64
		want  string
	}{
		{15000, "Rp 15.000"},
		{32782000, "Rp 32.782.000"},
		{0, "Rp 0"},
		{-50000, "Rp -50.000"},
		{1250, "Rp 1.250"}, // Sudah integer, nggak perlu rounding
	}

	for _, tt := range tests {
		got := FormatRupiah(tt.input)
		if got != tt.want {
			t.Errorf("FormatRupiah(%v) = %q, ingin %q", tt.input, got, tt.want)
		}
	}
}

func TestFormatRupiahCompact(t *testing.T) {
	tests := []struct {
		input int64
		want  string
	}{
		{1500000000, "Rp 1,5 M"},
		{2500000, "Rp 2,5 Jt"},
		{1500, "Rp 1,5 Rb"},
		{500, "Rp 500"},
		{1200000000, "Rp 1,2 M"},
	}

	for _, tt := range tests {
		got := FormatRupiahCompact(tt.input)
		if got != tt.want {
			t.Errorf("FormatRupiahCompact(%v) = %q, ingin %q", tt.input, got, tt.want)
		}
	}
}
