package main

import (
	"regexp"
	"strings"
)

// Remove qualquer pontuação da palavra
func sanitizeWord(w string) string {
	return strings.Trim(w, "!.,:;\"'")
}

// Adiciona [] ao redor da palavra ignorando pontuações
func surroundWord(word string) string {
	regex := regexp.MustCompile(`[^!.,:;"']+`)
	return regex.ReplaceAllString(word, `[$0]`)
}
