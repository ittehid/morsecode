package utils

import "fmt"

func WelcomeMessage() {
	fmt.Println("=Азбука Морзе=")
	fmt.Println("Что вы хотите сделать? Введите соответствующую цифру.")
	fmt.Println("1 - расшифровать код Морзе (вводите код через пробелы)")
	fmt.Println("2 - зашифровать сообщение")
}

func GetOption(optionNum *int) error {
	_, err := fmt.Scan(optionNum)
	return err
}
