package encryption

import (
	"bufio"
	"fmt"
	"morsecode/bdcode"
	"os"
	"strings"
)

func DecryptMorseMessage() {
	fmt.Println("Введите или вставьте код Морзе и в конце нажмите Enter")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		messageDecrypt := strings.TrimSpace(scanner.Text())
		messageDecrypt = strings.ReplaceAll(messageDecrypt, "  ", " * ")
		messageDecryptSlice := strings.Split(messageDecrypt, " ")
		fmt.Println(DecryptMorse(messageDecryptSlice))
	}
}

func DecryptMorse(morseCode []string) string {
	messageDecrypt := ""
	for _, val := range morseCode {
		if val == "*" {
			messageDecrypt += " "
		} else if en := decryptEN(val); en != "" {
			messageDecrypt += en
		} else if digits := decryptDigits(val); digits != "" {
			messageDecrypt += digits
		} else {
			return "Ошибка: код не найден"
		}
	}
	return messageDecrypt
}

func decryptEN(code string) string {
	for sim, cod := range bdcode.CodeEN {
		if cod == code {
			return sim
		}
	}
	return "Ошибка: символ не найден"
}

func decryptDigits(code string) string {
	for sim, cod := range bdcode.CodeDigits {
		if cod == code {
			return sim
		}
	}
	return "Ошибка: символ не найден"
}
