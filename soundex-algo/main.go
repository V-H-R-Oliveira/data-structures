package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const A = 97
const Z = 122

var CODES = map[byte]byte{
	'b': '1',
	'f': '1',
	'p': '1',
	'v': '1',
	'c': '2',
	'g': '2',
	'j': '2',
	'k': '2',
	'q': '2',
	's': '2',
	'x': '2',
	'z': '2',
	'd': '3',
	't': '3',
	'l': '4',
	'm': '5',
	'n': '5',
	'r': '6',
}

func soundex(text string) string {
	textBuffer := []byte(strings.ToLower(text))
	soundexBuffer := []byte{textBuffer[0]}

	for i, char := range textBuffer[1:] {
		currentCharCode, hasCode := CODES[char]

		if !hasCode {
			continue
		}

		previousChar := textBuffer[i]
		previousCharCode := CODES[previousChar]

		if previousCharCode == currentCharCode {
			continue
		}

		soundexBuffer = append(soundexBuffer, currentCharCode)
	}

	soundexBufferLength := len(soundexBuffer)

	if soundexBufferLength < 4 {
		padAmount := 4 - len(soundexBuffer)

		for i := 0; i < padAmount; i++ {
			soundexBuffer = append(soundexBuffer, '0')
		}
	} else if soundexBufferLength > 4 {
		soundexBuffer = soundexBuffer[:4]
	}

	return string(soundexBuffer)
}

func main() {
	// const s1 = "Rupert"
	// const s2 = "Robert"
	// const s3 = "Pfister"
	// const s4 = "Tymczak"

	// soundex(s1)
	// soundex(s2)
	// soundex(s3)
	// soundex(s4)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a string: ")
		read, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		if len(read) == 0 {
			continue
		}

		read = strings.Trim(read, "\n\r\t")

		fmt.Printf("Soundex(%s) = %s\n", read, soundex(read))

	}
}
