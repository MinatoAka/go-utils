package currency

import (
	"strconv"
	"strings"
)

// FormatRupiah mengubah angka int64 (satuan Rupiah) menjadi format Rupiah.
// Contoh: 15000000 -> "Rp 15.000.000"
func FormatRupiah(n int64) string {
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}
	
	s := strconv.FormatInt(n, 10)
	var b strings.Builder
	b.WriteString("Rp ")
	b.WriteString(sign)
	
	for i, c := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			b.WriteByte('.')
		}
		b.WriteRune(c)
	}
	return b.String()
}

// FormatRupiahCompact memformat angka besar menjadi singkat menggunakan pure integer math.
// Contoh: 1500000000 -> "Rp 1,5 M"
func FormatRupiahCompact(n int64) string {
	if n >= 1000000000 {
		return "Rp " + strconv.FormatInt(n/1000000000, 10) + "," + strconv.FormatInt((n%1000000000)/100000000, 10) + " M"
	}
	if n >= 1000000 {
		return "Rp " + strconv.FormatInt(n/1000000, 10) + "," + strconv.FormatInt((n%1000000)/100000, 10) + " Jt"
	}
	if n >= 1000 {
		return "Rp " + strconv.FormatInt(n/1000, 10) + "," + strconv.FormatInt((n%1000)/100, 10) + " Rb"
	}
	return FormatRupiah(n)
}

// FormatUSD mengubah angka float64 menjadi format USD dengan 2 desimal.
// Contoh: 1234567.89 -> "$1,234,567.89"
func FormatUSD(n float64) string {
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}
	
	// Format dengan 2 desimal
	val := strconv.FormatFloat(n, 'f', 2, 64)
	parts := strings.Split(val, ".")
	intPart := parts[0]
	decPart := parts[1]
	
	// Tambahkan koma pemisah ribuan
	var b strings.Builder
	b.WriteString("$")
	b.WriteString(sign)
	
	for i, c := range intPart {
		if i > 0 && (len(intPart)-i)%3 == 0 {
			b.WriteByte(',')
		}
		b.WriteRune(c)
	}
	
	b.WriteByte('.')
	b.WriteString(decPart)
	
	return b.String()
}

// FormatUSDCompact memformat angka USD besar menjadi singkat.
// Contoh: 1500000000 -> "$1.5B", 2500000 -> "$2.5M"
func FormatUSDCompact(n float64) string {
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}
	
	if n >= 1000000000 {
		val := strconv.FormatFloat(n/1000000000, 'f', 1, 64)
		return "$" + sign + val + "B"
	}
	if n >= 1000000 {
		val := strconv.FormatFloat(n/1000000, 'f', 1, 64)
		return "$" + sign + val + "M"
	}
	if n >= 1000 {
		val := strconv.FormatFloat(n/1000, 'f', 1, 64)
		return "$" + sign + val + "K"
	}
	return FormatUSD(n)
}
