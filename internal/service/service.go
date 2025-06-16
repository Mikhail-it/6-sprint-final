package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(in string) bool {
	clean := strings.TrimSpace(in)
	for _, char := range clean {
		if char != '.' && char != '-' && char != ' ' {
			return false
		}
	}
	return true
}

func Convert(in string) string {
	if isMorse(in) {
		return morse.ToText(in)
	} else {
		return morse.ToMorse(in)
	}
}
