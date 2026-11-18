# go-utils

[![Go Test](https://github.com/MinatoAka/go-utils/actions/workflows/go.yml/badge.svg)](https://github.com/MinatoAka/go-utils/actions/workflows/go.yml)
[![Go Version](https://img.shields.io/badge/Go-1.27.3-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Saweria](https://img.shields.io/badge/Support-Saweria-orange)](https://saweria.co/AKaAturDompet)

Kumpulan utilitas Go ringan untuk membantu tugas-tugas umum seperti format mata uang, manipulasi string, environment variable, JSON, dan format waktu.

## Fitur

- **currency** — format angka menjadi mata uang Rupiah biasa maupun ringkas
- **env** — membaca environment variable dan memuat file `.env`
- **json** — membaca dan menulis JSON secara pretty print dengan penulisan atomik
- **string** — konversi string ke angka dan sanitasi teks
- **timex** — format tanggal dan waktu untuk keperluan database maupun format Indonesia
- **Unit tests** — tersedia untuk setiap package
- **Tanpa dependency eksternal** — hanya menggunakan Go standard library

## Instalasi

```bash
go get github.com/MinatoAka/go-utils
```

Memerlukan Go 1.27.3 atau lebih baru.

## Penggunaan

### Currency

```go
package main

import (
    "fmt"

    "github.com/MinatoAka/go-utils/currency"
)

func main() {
    fmt.Println(currency.FormatRupiah(32782000))
    // Rp 32.782.000

    fmt.Println(currency.FormatRupiahCompact(2500000))
    // Rp 2,5 Jt
}
```

### Environment

```go
package main

import (
    "fmt"

    "github.com/MinatoAka/go-utils/env"
)

func main() {
    _ = env.LoadDotEnv(".env")

    port := env.GetEnv("APP_PORT", "8080")
    fmt.Println(port)
}
```

Contoh `.env`:

```env
APP_PORT=8080
DB_HOST=localhost
DB_PORT=5432
```

### JSON

```go
package main

import "github.com/MinatoAka/go-utils/json"

type Config struct {
    Name string `json:"name"`
    Port int    `json:"port"`
}

func main() {
    config := Config{Name: "server", Port: 8080}

    if err := json.WriteFile("config.json", config); err != nil {
        panic(err)
    }

    var loaded Config
    if err := json.ReadFile("config.json", &loaded); err != nil {
        panic(err)
    }
}
```

### String

```go
package main

import (
    "fmt"

    stringutil "github.com/MinatoAka/go-utils/string"
)

func main() {
    fmt.Println(stringutil.ToInt("42", 0))
    fmt.Println(stringutil.ToFloat64("12.5", 0))
    fmt.Println(stringutil.Sanitize("  GoLang  "))
    // 42
    // 12.5
    // golang
}
```

### Date and Time

```go
package main

import (
    "fmt"
    "time"

    "github.com/MinatoAka/go-utils/timex"
)

func main() {
    now := time.Now()

    fmt.Println(timex.FormatDateTime(now))
    fmt.Println(timex.FormatDateID(now))
    fmt.Println(timex.FormatTimeOnly(now))
    fmt.Println(timex.NowDateTime())
}
```

## Struktur Package

```text
currency/   Format mata uang Rupiah dan format ringkas
env/        Environment variable dan pembacaan file .env
json/       Pembacaan dan penulisan file JSON
string/     Konversi dan sanitasi string
timex/      Format tanggal dan waktu
```

## Menjalankan Test

```bash
go test ./...
```

Menjalankan pemeriksaan statis:

```bash
go vet ./...
```

## Lisensi

Project ini dilisensikan di bawah lisensi [MIT](LICENSE).
