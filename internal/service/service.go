package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(in string) bool {
	for i := 0; i < len(in); i++ {
		if in[i] != '.' && in[i] != '-' && in[i] != ' ' && in[i] != '\n' && in[i] != '\r' {
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
