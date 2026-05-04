package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(input string) bool {
	for _, symb := range input {
		if symb != '.' && symb != '-' && symb != ' ' {
			return false
		}
	}
	return true
}
func Convert(input string) (string, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return "", errors.New("Пустая строка")
	}

	if isMorse(input) {
		result := morse.ToText(input)
		if strings.TrimSpace(result) == "" {
			return "", errors.New("Ошибка перевода в текст")
		}
		return result, nil
	}

	result := morse.ToMorse(input)

	if result == "" {
		return "", errors.New("Ошибка перевода в Морзе")
	}

	return result, nil

}
