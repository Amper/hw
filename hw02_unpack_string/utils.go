package hw02unpackstring

import (
	"fmt"

	"github.com/pkg/errors"
)

var defaultLetterCount = 1

func flushLetter(s *unpackState, r *unpackResult, letter *rune, digit *int) error {
	err := r.WriteLetter(letter, digit)
	if err != nil {
		return newUnpackError(s, "processing error", err)
	}
	return s.SetLetter(nil)
}

func newUnpackError(s *unpackState, msg string, err ...error) error {
	if len(err) == 0 {
		return errors.Wrap(
			ErrInvalidString,
			fmt.Sprintf("there is %v in string \"%v\" on position %d", msg, s.Value(), s.Pos()),
		)
	}
	return errors.Wrap(
		ErrInvalidString,
		fmt.Sprintf("there is %v in string \"%v\" on position %d: %v", msg, s.Value(), s.Pos(), err[0]),
	)
}

func IsDigit(char rune) bool {
	return char != 0 && GetCharType(char) == CharTypeDigit
}

func IsSlash(char rune) bool {
	return char != 0 && GetCharType(char) == CharTypeSlash
}
