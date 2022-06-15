package hw02unpackstring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCharTypes(t *testing.T) {
	tests := []struct {
		input    rune
		expected CharType
	}{
		{input: '0', expected: CharTypeDigit},
		{input: '1', expected: CharTypeDigit},
		{input: '2', expected: CharTypeDigit},
		{input: '3', expected: CharTypeDigit},
		{input: '4', expected: CharTypeDigit},
		{input: '5', expected: CharTypeDigit},
		{input: '6', expected: CharTypeDigit},
		{input: '7', expected: CharTypeDigit},
		{input: '8', expected: CharTypeDigit},
		{input: '9', expected: CharTypeDigit},
		{input: 'a', expected: CharTypeLetter},
		{input: 'b', expected: CharTypeLetter},
		{input: 'd', expected: CharTypeLetter},
		{input: 'g', expected: CharTypeLetter},
		{input: 'o', expected: CharTypeLetter},
		{input: 'z', expected: CharTypeLetter},
		/*
			The supported alphabet is not clearly indicated in the task conditions.
			In addition, there are conflicting conditions regarding the support of "\n".
			Therefore, it is implemented this way, but to support any of the behaviors,
			it is enough to change only one place and tests.
		*/
		{input: '\n', expected: CharTypeLetter},
		{input: '\t', expected: CharTypeLetter},
		{input: '\\', expected: CharTypeSlash},
		{input: 0, expected: CharTypeInvalid},
		{input: '-', expected: CharTypeInvalid},
		{input: '.', expected: CharTypeInvalid},
		{input: '_', expected: CharTypeInvalid},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(string(tc.input), func(t *testing.T) {
			result := GetCharType(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}
