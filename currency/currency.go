package currency

import (
	"math"
	"strconv"
	"strings"
)

// FormatRupiah mengubah angka float64 menjadi format Rupiah (contoh: "Rp 15.000.000").
func FormatRupiah(n float64) string {
	val := int64(math.Round(n))
	sign := ""
	if val < 0 {
		sign = "-"
		val = -val
	}
	s := strconv.FormatInt(val, 10)
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

// FormatRupiahCompact memformat angka besar menjadi singkat (contoh: "15 Jt", "1,5 M").
func FormatRupiahCompact(n float64) string {
	if n >= 1000000000 {
		val := strings.Replace(strconv.FormatFloat(n/1000000000, 'f', 1, 64), ".", ",", 1)
		return "Rp " + val + " M"
	}
	if n >= 1000000 {
		val := strings.Replace(strconv.FormatFloat(n/1000000, 'f', 1, 64), ".", ",", 1)
		return "Rp " + val + " Jt"
	}
	if n >= 1000 {
		val := strings.Replace(strconv.FormatFloat(n/1000, 'f', 1, 64), ".", ",", 1)
		return "Rp " + val + " Rb"
	}
	return FormatRupiah(n)
}
