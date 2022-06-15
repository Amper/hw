package hw02unpackstring

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(value string) (string, error) {
	var (
		result = newUnpackResult()
		state  = newUnpackState(value)
		err    error
	)

	for state.HasNext() {
		state.Next()

		switch GetCharType(state.CurrChar()) {
		case CharTypeInvalid:
			err = newUnpackError(state, fmt.Sprintf("invalid character '%v'", state.CurrChar()))
		case CharTypeLetter:
			err = processLetter(state, result)
		case CharTypeSlash:
			err = processSlash(state, result)
		case CharTypeDigit:
			err = processDigit(state, result)
		}

		if err != nil {
			return "", err
		}
	}

	return result.String(), nil
}

func processLetter(state *unpackState, result *unpackResult) error {
	currChar, nextChar := state.CurrChar(), state.NextChar()
	if IsDigit(nextChar) {
		return state.SetLetter(&currChar)
	}
	return flushLetter(state, result, &currChar, &defaultLetterCount)
}

func processSlash(state *unpackState, result *unpackResult) error {
	prevChar, currChar, nextChar := state.Chars()
	letter := state.Letter()

	if IsSlash(prevChar) {
		if letter != nil {
			return state.SetLetter(nil)
		}
		if IsDigit(nextChar) {
			return state.SetLetter(&currChar)
		}

		err := result.WriteLetter(&currChar, &defaultLetterCount)
		if err != nil {
			return newUnpackError(state, "processing error", err)
		}
		if !IsSlash(nextChar) {
			return state.SetLetter(nil)
		}

		return state.SetLetter(&currChar)
	}

	if nextChar == 0 {
		return newUnpackError(state, "trailing slash")
	}

	return nil
}

func processDigit(state *unpackState, result *unpackResult) error {
	currChar := state.CurrChar()

	if letter := state.Letter(); letter != nil {
		digit, err := strconv.Atoi(string(currChar))
		if err != nil {
			return newUnpackError(state, "processing error", err)
		}
		return flushLetter(state, result, letter, &digit)
	}

	if !IsSlash(state.PrevChar()) {
		return newUnpackError(state, fmt.Sprintf("invalid character '%v'", currChar))
	}

	if IsDigit(state.NextChar()) {
		return state.SetLetter(&currChar)
	}

	return flushLetter(state, result, &currChar, &defaultLetterCount)
}
