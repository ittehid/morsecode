package encryption

import (
	"bufio"
	"fmt"
	"morsecode/bdcode"
	"os"
	"strings"
	"unicode"
)

func EncryptMorseMessage() {
	fmt.Println("Введите текст, через пробел и в конце нажмите Enter")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		messageEncrypt := scanner.Text()
		messageEncrypt = strings.TrimSpace(strings.ToUpper(messageEncrypt))
		messageEncrypt = strings.ReplaceAll(messageEncrypt, " ", "*")
		fmt.Println(checkStr(messageEncrypt))
	}
}

func checkStr(str string) string {
	var encryptedChar string
	strLatin := false
	strCyril := false

	for _, val := range str {
		if unicode.In(val, unicode.Latin) {
			encryptedChar += encryptEN(string(val)) + " "
			strLatin = true
		} else if unicode.In(val, unicode.Cyrillic) {
			strCyril = true
		} else if unicode.In(val, unicode.Digit) {
			encryptedChar += encryptDigits(string(val)) + " "
		} else if val == '*' {
			encryptedChar += " "
		}
	}

	if strLatin && strCyril {
		return "Ошибка: предложение содержит смешанные языки"
	}
	return encryptedChar
}

func encryptEN(strSim string) string {
	if cod, ok := bdcode.CodeEN[strSim]; ok {
		return cod
	}
	return "Ошибка: символ латиницы не найден"
}

func encryptDigits(strSim string) string {
	if cod, ok := bdcode.CodeDigits[strSim]; ok {
		return cod
	}
	return "Ошибка: цифра не найдена"
}
