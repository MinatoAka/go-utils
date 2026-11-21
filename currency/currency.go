package currency

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strconv"
	"strings"
)

const (
	// CurrencyIDR untuk Rupiah.
	CurrencyIDR = "IDR"
	// CurrencyUSD untuk Dollar.
	CurrencyUSD = "USD"
)

// ParseAmount mengonversi amount dari berbagai tipe ke float64.
// Mendukung numeric type, string, []byte, json.Number, pointer,
// fmt.Stringer, big.Float, big.Int, dan custom numeric type via reflect.
func ParseAmount(amount any, currency string) (float64, error) {
	return parseAmount(amount, currency)
}

// FormatMoneyE format angka jadi mata uang dan mengembalikan error jika input invalid.
func FormatMoneyE(amount any, currency string) (string, error) {
	num, err := parseAmount(amount, currency)
	if err != nil {
		return "", err
	}
	return formatMoneyFloat(num, currency), nil
}

// FormatMoney format angka jadi mata uang.
// Otomatis ngenalin "IDR", "RP", "RUPIAH" -> Rupiah
// dan "USD", "$", "DOLLAR", "US" -> Dollar.
//
// amount bisa:
// - int, int8, int16, int32, int64
// - uint, uint8, uint16, uint32, uint64, uintptr
// - float32, float64
// - string, []byte, json.Number
// - fmt.Stringer
// - *big.Float, big.Float, *big.Int, big.Int
// - pointer ke numeric type
// - custom numeric type
//
// Contoh string yang didukung:
// - "1500"
// - "1500.00"
// - "1,500.00"
// - "1.500,00"
// - "Rp 1.500.000"
// - "$1,234.56"
// - "(1,234.56)"
func FormatMoney(amount any, currency string) string {
	num, err := parseAmount(amount, currency)
	if err != nil {
		return ""
	}
	return formatMoneyFloat(num, currency)
}

// FormatMoneyCompactE format angka jadi versi ringkas dan mengembalikan error jika input invalid.
func FormatMoneyCompactE(amount any, currency string) (string, error) {
	num, err := parseAmount(amount, currency)
	if err != nil {
		return "", err
	}
	return formatMoneyCompactFloat(num, currency), nil
}

// FormatMoneyCompact format angka jadi versi ringkas.
//
// Contoh:
// FormatMoneyCompact(1500000, "IDR") -> "Rp 1,5 Jt"
// FormatMoneyCompact(1500, "USD")   -> "$1.5K"
func FormatMoneyCompact(amount any, currency string) string {
	num, err := parseAmount(amount, currency)
	if err != nil {
		return ""
	}
	return formatMoneyCompactFloat(num, currency)
}

// FormatRupiah shortcut buat FormatMoney(amount, "IDR").
func FormatRupiah(amount any) string {
	return FormatMoney(amount, CurrencyIDR)
}

// FormatRupiahCompact shortcut buat FormatMoneyCompact(amount, "IDR").
func FormatRupiahCompact(amount any) string {
	return FormatMoneyCompact(amount, CurrencyIDR)
}

// FormatUSD shortcut buat FormatMoney(amount, "USD").
func FormatUSD(amount any) string {
	return FormatMoney(amount, CurrencyUSD)
}

// FormatUSDCompact shortcut buat FormatMoneyCompact(amount, "USD").
func FormatUSDCompact(amount any) string {
	return FormatMoneyCompact(amount, CurrencyUSD)
}

func normalizeCurrency(currency string) string {
	c := strings.ToUpper(strings.TrimSpace(currency))
	c = strings.ReplaceAll(c, " ", "")

	switch c {
	case "IDR", "RP", "RUPIAH":
		return CurrencyIDR
	case "USD", "$", "DOLLAR", "DOLLARS", "US", "US$":
		return CurrencyUSD
	default:
		return c
	}
}

func stripCurrencySymbols(s string) string {
	s = strings.TrimSpace(s)

	tokens := []string{
		"Rp.",
		"Rp",
		"IDR",
		"RUPIAH",
		"$",
		"USD",
		"US$",
		"DOLLARS",
		"DOLLAR",
		"US",
	}

	for _, token := range tokens {
		for {
			trimmed := strings.TrimSpace(s)
			lower := strings.ToLower(trimmed)
			tokenLower := strings.ToLower(token)
			changed := false

			if strings.HasPrefix(lower, tokenLower) {
				s = strings.TrimSpace(trimmed[len(token):])
				changed = true
			}

			if strings.HasSuffix(strings.ToLower(s), tokenLower) {
				s = strings.TrimSpace(s[:len(s)-len(token)])
				changed = true
			}

			if !changed {
				break
			}
		}
	}

	// Bersihkan sisa simbol yang sering menempel.
	s = strings.ReplaceAll(s, "Rp", "")
	s = strings.ReplaceAll(s, "rp", "")
	s = strings.ReplaceAll(s, "$", "")
	s = strings.TrimSpace(s)

	return s
}

func parseNumericString(s, currency string) (float64, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, fmt.Errorf("amount string kosong")
	}

	s = stripCurrencySymbols(s)
	s = strings.TrimSpace(s)

	// Hapus whitespace biasa dan whitespace unicode yang sering muncul dari copy-paste.
	for _, ws := range []string{" ", "\t", "\n", "\r", "\u00A0", "\u202F", "\u2009"} {
		s = strings.ReplaceAll(s, ws, "")
	}

	negative := false

	if strings.HasPrefix(s, "-") {
		negative = true
		s = s[1:]
	} else if strings.HasPrefix(s, "+") {
		s = s[1:]
	}

	// Dukungan format akuntansi: (1,234.56) = -1234.56
	if strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")") {
		negative = true
		s = s[1 : len(s)-1]
	}

	if s == "" {
		return 0, fmt.Errorf("amount tidak valid")
	}

	dot := strings.LastIndex(s, ".")
	comma := strings.LastIndex(s, ",")

	var decimalChar rune
	var groupingChar rune

	switch {
	case dot >= 0 && comma >= 0:
		// Kalau ada dua-duanya, pemisah terakhir dianggap desimal.
		if dot > comma {
			decimalChar = '.'
			groupingChar = ','
		} else {
			decimalChar = ','
			groupingChar = '.'
		}

	case dot >= 0:
		dots := strings.Count(s, ".")
		fracLen := len(s) - dot - 1

		// Aturan pragmatis:
		// - multiple dot = grouping
		// - satu dot dengan 3 digit di belakang = grouping
		// - selain itu = decimal
		if dots > 1 || fracLen == 3 {
			groupingChar = '.'
		} else {
			decimalChar = '.'
		}

	case comma >= 0:
		commas := strings.Count(s, ",")
		fracLen := len(s) - comma - 1

		// Aturan pragmatis:
		// - multiple comma = grouping
		// - satu comma dengan 3 digit di belakang = grouping
		// - selain itu = decimal
		if commas > 1 || fracLen == 3 {
			groupingChar = ','
		} else {
			decimalChar = ','
		}
	}

	if groupingChar != 0 {
		s = strings.ReplaceAll(s, string(groupingChar), "")
	}

	if decimalChar != 0 {
		s = strings.Replace(s, string(decimalChar), ".", 1)
	}

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("amount tidak valid: %w", err)
	}

	if negative {
		f = -f
	}

	return f, nil
}

func parseAmount(amount any, currency string) (float64, error) {
	switch v := amount.(type) {
	case nil:
		return 0, fmt.Errorf("amount nil")

	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil

	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case uintptr:
		return float64(v), nil

	case float32:
		f := float64(v)
		if math.IsNaN(f) || math.IsInf(f, 0) {
			return 0, fmt.Errorf("amount float tidak valid")
		}
		return f, nil

	case float64:
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return 0, fmt.Errorf("amount float tidak valid")
		}
		return v, nil

	case json.Number:
		return strconv.ParseFloat(string(v), 64)

	case string:
		return parseNumericString(v, currency)

	case []byte:
		return parseNumericString(string(v), currency)

	case *big.Float:
		if v == nil {
			return 0, fmt.Errorf("big.Float nil")
		}
		f, _ := v.Float64()
		if math.IsNaN(f) || math.IsInf(f, 0) {
			return 0, fmt.Errorf("big.Float tidak valid")
		}
		return f, nil

	case big.Float:
		f, _ := v.Float64()
		if math.IsNaN(f) || math.IsInf(f, 0) {
			return 0, fmt.Errorf("big.Float tidak valid")
		}
		return f, nil

	case *big.Int:
		if v == nil {
			return 0, fmt.Errorf("big.Int nil")
		}
		return strconv.ParseFloat(v.String(), 64)

	case big.Int:
		return strconv.ParseFloat(v.String(), 64)

	case fmt.Stringer:
		return parseNumericString(v.String(), currency)

	default:
		return parseAmountByReflect(amount, currency)
	}
}

func parseAmountByReflect(amount any, currency string) (float64, error) {
	rv := reflect.ValueOf(amount)
	if !rv.IsValid() {
		return 0, fmt.Errorf("amount invalid")
	}

	// Dereference pointer/interface.
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		if rv.IsNil() {
			return 0, fmt.Errorf("amount nil pointer/interface")
		}
		rv = rv.Elem()
	}

	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(rv.Int()), nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return float64(rv.Uint()), nil

	case reflect.Float32, reflect.Float64:
		f := rv.Float()
		if math.IsNaN(f) || math.IsInf(f, 0) {
			return 0, fmt.Errorf("amount float tidak valid")
		}
		return f, nil

	case reflect.String:
		return parseNumericString(rv.String(), currency)

	case reflect.Slice:
		if rv.Type().Elem().Kind() == reflect.Uint8 {
			return parseNumericString(string(rv.Bytes()), currency)
		}
	}

	if rv.CanInterface() {
		// Last resort: render pakai %v lalu parse sebagai string.
		// Ini membantu tipe numerik custom yang tidak tertangkap switch di atas.
		return parseNumericString(fmt.Sprintf("%v", rv.Interface()), currency)
	}

	return 0, fmt.Errorf("tipe amount tidak didukung: %T", amount)
}

func formatMoneyFloat(num float64, currency string) string {
	sign := ""
	if num < 0 {
		sign = "-"
		num = -num
	}

	switch normalizeCurrency(currency) {
	case CurrencyIDR:
		// IDR tanpa desimal.
		s := strconv.FormatFloat(num, 'f', 0, 64)
		return sign + "Rp " + groupThousands(s, '.')

	case CurrencyUSD:
		// USD selalu 2 desimal.
		val := strconv.FormatFloat(num, 'f', 2, 64)
		parts := strings.Split(val, ".")
		if len(parts) != 2 {
			return sign + "$" + groupThousands(val, ',')
		}
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

func formatMoneyCompactFloat(num float64, currency string) string {
	sign := ""
	if num < 0 {
		sign = "-"
		num = -num
	}

	switch normalizeCurrency(currency) {
	case CurrencyIDR:
		// Gunakan int64 selama masih aman supaya konsisten dengan perilaku truncation lama.
		if num <= float64(math.MaxInt64) {
			n := int64(num)

			switch {
			case n >= 1_000_000_000:
				return sign + "Rp " +
					strconv.FormatInt(n/1_000_000_000, 10) + "," +
					strconv.FormatInt((n%1_000_000_000)/100_000_000, 10) +
					" M"

			case n >= 1_000_000:
				return sign + "Rp " +
					strconv.FormatInt(n/1_000_000, 10) + "," +
					strconv.FormatInt((n%1_000_000)/100_000, 10) +
					" Jt"

			case n >= 1_000:
				return sign + "Rp " +
					strconv.FormatInt(n/1_000, 10) + "," +
					strconv.FormatInt((n%1_000)/100, 10) +
					" Rb"

			default:
				return sign + formatMoneyFloat(float64(n), CurrencyIDR)
			}
		}

		// Fallback untuk angka sangat besar di luar jangkauan int64.
		switch {
		case num >= 1_000_000_000:
			v := num / 1_000_000_000
			return sign + "Rp " + strings.Replace(strconv.FormatFloat(v, 'f', 1, 64), ".", ",", 1) + " M"

		case num >= 1_000_000:
			v := num / 1_000_000
			return sign + "Rp " + strings.Replace(strconv.FormatFloat(v, 'f', 1, 64), ".", ",", 1) + " Jt"

		case num >= 1_000:
			v := num / 1_000
			return sign + "Rp " + strings.Replace(strconv.FormatFloat(v, 'f', 1, 64), ".", ",", 1) + " Rb"

		default:
			return sign + formatMoneyFloat(num, CurrencyIDR)
		}

	case CurrencyUSD:
		switch {
		case num >= 1_000_000_000:
			return sign + "$" + strconv.FormatFloat(num/1_000_000_000, 'f', 1, 64) + "B"

		case num >= 1_000_000:
			return sign + "$" + strconv.FormatFloat(num/1_000_000, 'f', 1, 64) + "M"

		case num >= 1_000:
			return sign + "$" + strconv.FormatFloat(num/1_000, 'f', 1, 64) + "K"

		default:
			return sign + formatMoneyFloat(num, CurrencyUSD)
		}

	default:
		return ""
	}
}
