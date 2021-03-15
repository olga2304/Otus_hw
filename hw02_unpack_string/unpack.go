package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if _, err := strconv.Atoi(s); err == nil {
		return "", errors.New("invalid string")
	}

	var a rune
	b := true
	var str strings.Builder

	for _, char := range s {
		if unicode.IsDigit(char) && b {
			m := int(char - '0')
			if m == 0 {
				l := str
				str.Reset()
				str.WriteString(strings.Repeat(string(a), l.Len()-1))
			} else {
				r := strings.Repeat(string(a), m-1)
				str.WriteString(r)
			}
		} else {
			if b {
				str.WriteRune(char)
			}
			a = char
		}
	}

	return str.String(), nil
}
