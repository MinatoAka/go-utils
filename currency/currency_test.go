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
		{1250, "Rp 1.250"},
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

func TestFormatUSD(t *testing.T) {
	tests := []struct {
		input float64
		want  string
	}{
		{1234567.89, "$1,234,567.89"},
		{0, "$0.00"},
		{-500.5, "$-500.50"},
		{999999.99, "$999,999.99"},
		{1000, "$1,000.00"},
	}

	for _, tt := range tests {
		got := FormatUSD(tt.input)
		if got != tt.want {
			t.Errorf("FormatUSD(%v) = %q, ingin %q", tt.input, got, tt.want)
		}
	}
}

func TestFormatUSDCompact(t *testing.T) {
	tests := []struct {
		input float64
		want  string
	}{
		{1500000000, "$1.5B"},
		{2500000, "$2.5M"},
		{1500, "$1.5K"},
		{500, "$500.00"},
		{-1200000, "$-1.2M"},
	}

	for _, tt := range tests {
		got := FormatUSDCompact(tt.input)
		if got != tt.want {
			t.Errorf("FormatUSDCompact(%v) = %q, ingin %q", tt.input, got, tt.want)
		}
	}
}
