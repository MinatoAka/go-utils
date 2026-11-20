package currency

import "testing"

func TestFormatMoney(t *testing.T) {
	tests := []struct {
		input    any
		currency string
		want     string
	}{
		{1500000, "IDR", "Rp 1.500.000"},
		{1500000, "rp", "Rp 1.500.000"},
		{1500000, "Rupiah", "Rp 1.500.000"},
		{1500.5, "USD", "$1,500.50"},
		{1500.5, "$", "$1,500.50"},
		{1500.5, "dollar", "$1,500.50"},
		{-2500000, "IDR", "-Rp 2.500.000"},
		{-99.9, "USD", "-$99.90"},
		{0, "IDR", "Rp 0"},
		{0, "USD", "$0.00"},
	}

	for _, tt := range tests {
		got := FormatMoney(tt.input, tt.currency)
		if got != tt.want {
			t.Errorf("FormatMoney(%v, %q) = %q, ingin %q", tt.input, tt.currency, got, tt.want)
		}
	}
}

func TestFormatMoneyCompact(t *testing.T) {
	tests := []struct {
		input    any
		currency string
		want     string
	}{
		{int64(1500), "IDR", "Rp 1,5 Rb"},
		{int64(1500), "rp", "Rp 1,5 Rb"},
		{int64(1500000), "IDR", "Rp 1,5 Jt"},
		{int64(2500000000), "IDR", "Rp 2,5 M"},
		{int64(500), "IDR", "Rp 500"},
		{int64(-1500000), "IDR", "-Rp 1,5 Jt"},
		{1500.0, "USD", "$1.5K"},
		{1500.0, "$", "$1.5K"},
		{1500000.0, "USD", "$1.5M"},
		{2500000000.0, "USD", "$2.5B"},
		{500.0, "USD", "$500.00"},
		{-1500.0, "USD", "-$1.5K"},
		{0, "IDR", "Rp 0"},
		{0.0, "USD", "$0.00"},
	}

	for _, tt := range tests {
		got := FormatMoneyCompact(tt.input, tt.currency)
		if got != tt.want {
			t.Errorf("FormatMoneyCompact(%v, %q) = %q, ingin %q", tt.input, tt.currency, got, tt.want)
		}
	}
}

func TestFormatRupiah(t *testing.T) {
	tests := []struct {
		input int64
		want  string
	}{
		{15000, "Rp 15.000"},
		{32782000, "Rp 32.782.000"},
		{0, "Rp 0"},
		{-50000, "-Rp 50.000"},
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
		{1500, "Rp 1,5 Rb"},
		{15000, "Rp 15,0 Rb"},
		{1250, "Rp 1,2 Rb"},
		{1500000, "Rp 1,5 Jt"},
		{32782000, "Rp 32,7 Jt"},
		{2500000000, "Rp 2,5 M"},
		{1000, "Rp 1,0 Rb"},
		{1000000, "Rp 1,0 Jt"},
		{1000000000, "Rp 1,0 M"},
		{500, "Rp 500"},
		{0, "Rp 0"},
		{-50000, "-Rp 50,0 Rb"},
		{-1500000, "-Rp 1,5 Jt"},
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
		{15000, "$15,000.00"},
		{1500.5, "$1,500.50"},
		{0, "$0.00"},
		{-50000, "-$50,000.00"},
		{99.9, "$99.90"},
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
		{1500, "$1.5K"},
		{15000, "$15.0K"},
		{1500000, "$1.5M"},
		{2500000000, "$2.5B"},
		{1000, "$1.0K"},
		{1000000, "$1.0M"},
		{1000000000, "$1.0B"},
		{500, "$500.00"},
		{999.99, "$999.99"},
		{0, "$0.00"},
		{-1500, "-$1.5K"},
		{-50000, "-$50.0K"},
	}

	for _, tt := range tests {
		got := FormatUSDCompact(tt.input)
		if got != tt.want {
			t.Errorf("FormatUSDCompact(%v) = %q, ingin %q", tt.input, got, tt.want)
		}
	}
}
}
