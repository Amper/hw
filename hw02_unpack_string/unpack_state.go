package hw02unpackstring

type unpackState struct {
	*LookaheadIterator
	letter *rune
	err    error
}

func newUnpackState(value string) *unpackState {
	return &unpackState{
		LookaheadIterator: NewLookaheadIterator(value),
		letter:            nil,
		err:               nil,
	}
}

func (s *unpackState) HasLetter() bool {
	return s.letter != nil && *s.letter != 0
}

func (s *unpackState) Letter() *rune {
	return s.letter
}

func (s *unpackState) SetLetter(letter *rune) error {
	s.letter = letter
	return nil
}
