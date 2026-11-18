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
