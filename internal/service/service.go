package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// isMorse определяет, является ли строка кодом Морзе.
// Примерная эвристика — если строка содержит только точки, тире, пробелы и/или специальные разделители.
func isMorse(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" {
		return false
	}
	for _, r := range s {
		if !(r == '.' || r == '-' || unicode.IsSpace(r)) {
			return false
		}
	}
	return true
}

// ConvertAutoDetect автоматически определяет тип входа и конвертирует его
func ConvertAutoDetect(input string) (string, error) {
	trimmed := strings.TrimSpace(input)

	if trimmed == "" {
		return "", errors.New("input is empty")
	}

	if isMorse(trimmed) {
		// Морзе -> Текст
		return morse.ToText(trimmed), nil
	}

	// Текст -> Морзе
	return morse.ToMorse(trimmed), nil
}
