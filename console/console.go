package console

import (
	"fmt"
	"time"
	"unicode"
)

// FmtTypeWriter mencetak teks dengan efek mengetik per karakter.
// Spasi dicetak lebih cepat (speed / 2) agar terlihat lebih alami.
// Secara otomatis menambahkan newline di akhir jika teks belum memilikinya.
func FmtTypeWriter(text string, speed time.Duration) {
	chars := []rune(text)
	for i, char := range chars {
		fmt.Printf("%c", char)
		
		if i == len(chars)-1 {
			continue
		}
		if unicode.IsSpace(char) {
			time.Sleep(speed / 2)
		} else {
			time.Sleep(speed)
		}
	}
	
	// Memastikan selalu ada newline di akhir agar prompt terminal berikutnya rapi
	if len(chars) == 0 || chars[len(chars)-1] != '\n' {
		fmt.Println()
	}
}

// FmtLogWriter mencetak teks dengan efek mengetik, diawali dengan timestamp.
// Format timestamp: "2006/01/02 15:04:05 "
func FmtLogWriter(text string, speed time.Duration) {
	prefix := time.Now().Format("2006/01/02 15:04:05 ")
	FmtTypeWriter(prefix+text, speed)
}
