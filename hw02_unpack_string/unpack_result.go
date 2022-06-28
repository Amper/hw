package hw02unpackstring

import "strings"

type unpackResult struct {
	strings.Builder
}

func newUnpackResult() *unpackResult {
	return &unpackResult{
		Builder: strings.Builder{},
	}
}

func (r *unpackResult) WriteLetter(letter *rune, digit *int) error {
	if letter != nil && digit != nil {
		if *digit > 0 {
			r.Grow(*digit)
			for i := 0; i < *digit; i++ {
				_, err := r.WriteRune(*letter)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
