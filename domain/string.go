package domain

import (
	"errors"
	"unicode/utf8"
)

var (
	ErrInvalidStringLength = errors.New("domain: invalid string length")
)

type String string

func (s String) String() string {
	return string(s)
}

func (s String) ValidateLength(min, max int) error {
	lenS := utf8.RuneCountInString(string(s))
	if lenS >= min && lenS <= max {
		return nil
	} else {
		return ErrInvalidStringLength
	}
}
