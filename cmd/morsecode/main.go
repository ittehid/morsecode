package main

import (
	"fmt"
	"morsecode/encryption"
	"morsecode/utils"
)

func main() {
	utils.WelcomeMessage()
	var optionNum int

	// Получаем введенный пользователем выбор (1 или 2)
	if err := utils.GetOption(&optionNum); err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	switch optionNum {
	case 1:
		encryption.DecryptMorseMessage()
	case 2:
		encryption.EncryptMorseMessage()
	default:
		fmt.Println("Ошибка ввода. Введите цифру 1 или 2")
	}
}
