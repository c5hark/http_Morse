package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func TranslateToMorse(text string) (string, error) {
	text = strings.TrimSpace(text)
	if text == "" {
		return "", errors.New("input text cannot be empty")
	}

	if isMorse(text) {
		result := morse.ToText(text)
		if result == "" {
			return "", errors.New("invalid morse code")
		}
		return result, nil
	}

	result := morse.ToMorse(text)
	if result == "" {
		return "", errors.New("invalid text")
	}
	return result, nil
}

func isMorse(text string) bool {
	for _, r := range text {
		switch r {
		case ' ', '-', '.', '/':
			continue
		default:
			return false
		}
	}
	return true
}
