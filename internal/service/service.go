// package service

package service

import (
	"errors"
	"strings"

	"github.com/es-x/6f/pkg/morse"
)

func Converter(s string) (string, error) {

	f := func(r rune) bool {
		return r != '.' && r != '-' && r != ' '
	}
	if s == "" {
		return "", errors.New("empty string")
	}
	var result string
	isMorse := strings.ContainsFunc(s, f)
	if isMorse {
		result = morse.ToMorse(s)
	} else {
		result = morse.ToText(s)
	}
	return result, nil
}
