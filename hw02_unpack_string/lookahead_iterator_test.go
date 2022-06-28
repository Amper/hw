package hw02unpackstring

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIterator(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "", expected: ""},
		{input: "a", expected: "0_a_"},
		{input: "ab", expected: "0_ab | 1ab_"},
		{input: "abc", expected: "0_ab | 1abc | 2bc_"},
		{input: "aaaa", expected: "0_aa | 1aaa | 2aaa | 3aa_"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result := getIteratorOutput(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}

func getIteratorOutput(value string) string {
	var (
		iterator  = NewLookaheadIterator(value)
		result    = make([]string, 0)
		runeValue = func(r rune) string {
			if r == 0 {
				return "_"
			}
			return string(r)
		}
	)

	for iterator.HasNext() {
		prev, cur, next := iterator.Next()
		result = append(result, fmt.Sprintf("%v%v%v%v", iterator.Pos(), runeValue(prev), runeValue(cur), runeValue(next)))
	}

	return strings.Join(result, " | ")
}
