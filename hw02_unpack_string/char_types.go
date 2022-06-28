package hw02unpackstring

import "unicode"

type CharType string

const (
	CharTypeLetter  = CharType("letter")
	CharTypeDigit   = CharType("digit")
	CharTypeSlash   = CharType("slash")
	CharTypeInvalid = CharType("invalid")
)

func GetCharType(char rune) CharType {
	switch {
	case unicode.IsDigit(char):
		return CharTypeDigit
	case unicode.IsLetter(char):
		return CharTypeLetter
	/*
		The supported alphabet is not clearly indicated in the task conditions..
		Therefore, it is implemented this way, but to support any of the behaviors,
		it is enough to change only one place and tests.
	*/
	case unicode.IsSpace(char):
		return CharTypeLetter
	case char == '\\':
		return CharTypeSlash
	default:
		return CharTypeInvalid
	}
}
