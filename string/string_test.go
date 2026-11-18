package string

import "testing"

func TestToFloat64(t *testing.T) {
	tests := []struct {
		input string
		def   float64
		want  float64
	}{
		{"12.5", 0.0, 12.5},
		{"  100  ", 0.0, 100.0},
		{"bukan_angka", 99.9, 99.9},
		{"", 5.0, 5.0},
	}

	for _, tt := range tests {
		got := ToFloat64(tt.input, tt.def)
		if got != tt.want {
			t.Errorf("ToFloat64(%q) = %v, ingin %v", tt.input, got, tt.want)
		}
	}
}

func TestToInt(t *testing.T) {
	tests := []struct {
		input string
		def   int
		want  int
	}{
		{"42", 0, 42},
		{"  7  ", 0, 7},
		{"abc", -1, -1},
	}

	for _, tt := range tests {
		got := ToInt(tt.input, tt.def)
		if got != tt.want {
			t.Errorf("ToInt(%q) = %v, ingin %v", tt.input, got, tt.want)
		}
	}
}

func TestSanitize(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"  HELLO  ", "hello"},
		{"GoLang", "golang"},
		{"", ""},
	}

	for _, tt := range tests {
		got := Sanitize(tt.input)
		if got != tt.want {
			t.Errorf("Sanitize(%q) = %q, ingin %q", tt.input, got, tt.want)
		}
	}
}
