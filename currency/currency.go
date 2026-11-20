package currency

import (
	"strconv"
	"strings"
)

const (
	// CurrencyIDR untuk Rupiah.
	CurrencyIDR = "IDR"
	// CurrencyUSD untuk Dollar.
	CurrencyUSD = "USD"
)

// FormatMoney format angka jadi mata uang.
// Otomatis ngenalin "IDR", "RP", "RUPIAH" -> Rupiah
// dan "USD", "$", "DOLLAR" -> Dollar.
func FormatMoney(amount any, currency string) string {
	var num float64
	switch v := amount.(type) {
	case int:
		num = float64(v)
	case int32:
		num = float64(v)
	case int64:
		num = float64(v)
	case float32:
		num = float64(v)
	case float64:
		num = v
	default:
		return ""
	}

	sign := ""
	if num < 0 {
		sign = "-"
		num = -num
	}

	switch strings.ToUpper(strings.TrimSpace(currency)) {
	case "IDR", "RP", "RUPIAH":
		// IDR tanpa desimal
		s := strconv.FormatFloat(num, 'f', 0, 64)
		return sign + "Rp " + groupThousands(s, '.')

	case "USD", "$", "DOLLAR", "US":
		// USD selalu 2 desimal
		val := strconv.FormatFloat(num, 'f', 2, 64)
		parts := strings.Split(val, ".")
		return sign + "$" + groupThousands(parts[0], ',') + "." + parts[1]

	default:
		return sign + strconv.FormatFloat(num, 'f', 2, 64)
	}
}

func groupThousands(s string, sep byte) string {
	var b strings.Builder
	for i, c := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			b.WriteByte(sep)
		}
		b.WriteRune(c)
	}
	return b.String()
}

// FormatMoneyCompact format angka jadi versi ringkas.
// FormatMoneyCompact(1500000, "IDR") -> "Rp 1,5 Jt"
// FormatMoneyCompact(1500, "USD") -> "$1.5K"
func FormatMoneyCompact(amount any, currency string) string {
	var num float64
	switch v := amount.(type) {
	case int:
		num = float64(v)
	case int32:
		num = float64(v)
	case int64:
		num = float64(v)
	case float32:
		num = float64(v)
	case float64:
		num = v
	default:
		return ""
	}

	sign := ""
	if num < 0 {
		sign = "-"
		num = -num
	}

	switch strings.ToUpper(strings.TrimSpace(currency)) {
	case "IDR", "RP", "RUPIAH":
		n := int64(num)
		if n >= 1000000000 {
			return sign + "Rp " + strconv.FormatInt(n/1000000000, 10) + "," + strconv.FormatInt((n%1000000000)/100000000, 10) + " M"
		}
		if n >= 1000000 {
			return sign + "Rp " + strconv.FormatInt(n/1000000, 10) + "," + strconv.FormatInt((n%1000000)/100000, 10) + " Jt"
		}
		if n >= 1000 {
			return sign + "Rp " + strconv.FormatInt(n/1000, 10) + "," + strconv.FormatInt((n%1000)/100, 10) + " Rb"
		}
		return sign + FormatMoney(num, CurrencyIDR)

	case "USD", "$", "DOLLAR", "US":
		if num >= 1000000000 {
			return sign + "$" + strconv.FormatFloat(num/1000000000, 'f', 1, 64) + "B"
		}
		if num >= 1000000 {
			return sign + "$" + strconv.FormatFloat(num/1000000, 'f', 1, 64) + "M"
		}
		if num >= 1000 {
			return sign + "$" + strconv.FormatFloat(num/1000, 'f', 1, 64) + "K"
		}
		return sign + FormatMoney(num, CurrencyUSD)

	default:
		return ""
	}
}

// FormatRupiah shortcut buat FormatMoney(n, "IDR").
func FormatRupiah(n int64) string {
	return FormatMoney(n, CurrencyIDR)
}

// FormatRupiahCompact shortcut buat FormatMoneyCompact(n, "IDR").
func FormatRupiahCompact(n int64) string {
	return FormatMoneyCompact(n, CurrencyIDR)
}

// FormatUSD shortcut buat FormatMoney(n, "USD").
func FormatUSD(n float64) string {
	return FormatMoney(n, CurrencyUSD)
}

// FormatUSDCompact shortcut buat FormatMoneyCompact(n, "USD").
func FormatUSDCompact(n float64) string {
	return FormatMoneyCompact(n, CurrencyUSD)
}
