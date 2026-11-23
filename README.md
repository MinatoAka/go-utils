# go-utils

[![Go Test](https://github.com/MinatoAka/go-utils/actions/workflows/go.yml/badge.svg)](https://github.com/MinatoAka/go-utils/actions/workflows/go.yml)
[![Go Version](https://img.shields.io/badge/Go-1.27.3-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Saweria](https://img.shields.io/badge/Support-Saweria-orange)](https://saweria.co/AKaAturDompet)

Kumpulan utilitas Go ringan untuk membantu tugas-tugas umum seperti format mata uang, manipulasi string, environment variable, JSON, console, dan format waktu.

## Fitur

- **console** — mencetak teks dengan efek mengetik dan timestamp untuk log
- **currency** — format angka menjadi mata uang `Rupiah` & `USD` biasa maupun ringkas
- **env** — membaca environment variable dan memuat file `.env`
- **json** — membaca dan menulis JSON
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

### Console

```go
package main

import (
    "time"

    "github.com/MinatoAka/go-utils/console"
)

func main() {
    console.FmtTypeWriter("Memproses permintaan...", 30*time.Millisecond)
    console.FmtLogWriter("Server berhasil dijalankan", 20*time.Millisecond)
}
```

- `FmtTypeWriter` mencetak teks per karakter dengan jeda spasi yang lebih cepat.
- `FmtLogWriter` menambahkan timestamp dengan format `2006/01/02 15:04:05` sebelum teks.

### Currency

```go
package main

import (
    "fmt"

    "github.com/MinatoAka/go-utils/currency"
)

func main() {
    // Format Rupiah
    fmt.Println(currency.FormatRupiah(32782000))
    // Rp 32.782.000

    fmt.Println(currency.FormatRupiahCompact(2500000))
    // Rp 2,5 Jt

    // Format USD
    fmt.Println(currency.FormatUSD(1234567.89))
    // $1,234,567.89

    fmt.Println(currency.FormatUSDCompact(2500000))
    // $2.5M
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

### 📦 JSON Package (`go-utils/json`)

A high-performance, type-safe JSON mini-database. Replaces traditional `os.ReadFile` and `os.WriteFile` with atomic writes, generics, and zero-boilerplate state management.

#### ✨ Features
- **Generics (Type-Safe):** No more `interface{}` casting. `Read[T]` and `Write[T]` return exactly what you expect.
- **Atomic Write:** Uses temporary files + `os.Rename` to prevent file corruption during power loss.
- **Read-Modify-Write:** The `Update` function safely handles reading, mutating, and saving data in one step.
- **Fast I/O:** Skips `fsync` by default for maximum execution speed (ideal for trading bots).
- **Pretty Print (Optional):** Built-in helpers for human-readable formatting when debugging.

#### 🚀 Usage

##### 1. Basic Read & Write
```go
package main

import (
    "fmt"
    "github.com/MinatoAka/go-utils/json"
)

type Config struct {
    Name string `json:"name"`
    Port int    `json:"port"`
}

func main() {
    cfg := Config{Name: "server", Port: 8080}

    // Write to file (Atomic & Fast)
    if err := json.Write("config.json", cfg); err != nil {
        panic(err)
    }

    // Read from file into a specific type
    loaded, err := json.Read[Config]("config.json")
    if err != nil {
        panic(err)
    }

    fmt.Printf("Loaded: %s on port %d\n", loaded.Name, loaded.Port)
}
```

##### 2. Load or Create (Initialization)
```go
defaultConfig := Config{Name: "default", Port: 3000}

// Reads if exists, creates with defaultConfig if missing
cfg, err := json.LoadOrCreate("config.json", defaultConfig)
if err != nil {
    panic(err)
}
```

##### 3. Update (Read-Modify-Write)
```go
err := json.Update("config.json", func(c *Config) {
    c.Port = 9090 // Only change the port, leave other fields intact
})
```

##### 4. Pretty Print (Debugging)
```go
// 1. Print to terminal
prettyStr, _ := json.MarshalPretty(cfg)
fmt.Println(prettyStr)

// 2. Write pretty JSON to file
json.WritePretty("debug.json", cfg)
```


## Struktur Package

```text
console/    Efek mengetik dan log dengan timestamp
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
