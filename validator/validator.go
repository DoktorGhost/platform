package validator

import (
	"fmt"
	"strings"
	"unicode"
)

// вспомогалтельные функции
func valid(name string) bool {
	for i, char := range name {
		// Проверяем, что символ является буквой, дефисом или пробелом
		if !unicode.IsLetter(char) && char != '-' && char != ' ' {
			return false
		}

		// Проверяем, что дефис или пробел не в начале или в конце строки
		if (char == '-' || char == ' ') && (i == 0 || i == len(name)-1) {
			return false
		}
	}
	return true
}

func Validator(name, surname, patronymic string) (string, error) {
	if !valid(name) {
		return "", fmt.Errorf("имя не должно содержать цифры или символы")
	}
	if !valid(surname) {
		return "", fmt.Errorf("фамилия не должна содержать цифры или символы")
	}
	if !valid(patronymic) {
		return "", fmt.Errorf("отчество не должно содержать цифры или символы")
	}
	fullName := strings.TrimSpace(name + " " + surname + " " + patronymic)
	return fullName, nil
}
