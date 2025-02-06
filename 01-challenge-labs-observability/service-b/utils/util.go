package utils

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func RemoveAccents(input string) string {
	t := norm.NFD.String(input)
	return strings.Map(func(r rune) rune {
		if unicode.Is(unicode.Mn, r) {
			return -1
		}
		return r
	}, t)
}
