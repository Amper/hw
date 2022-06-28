package hw02unpackstring

type LookaheadIterator struct {
	value    string
	runes    []rune
	position int
}

func NewLookaheadIterator(value string) *LookaheadIterator {
	return &LookaheadIterator{
		value:    value,
		runes:    []rune(value),
		position: -1,
	}
}

func (i *LookaheadIterator) HasNext() bool {
	return len(i.runes) > 0 && i.position < len(i.runes)-1
}

func (i *LookaheadIterator) Next() (prev rune, curr rune, next rune) {
	i.position++
	return i.Chars()
}

func (i *LookaheadIterator) Value() string {
	return i.value
}

func (i *LookaheadIterator) Pos() int {
	return i.position
}

func (i *LookaheadIterator) PrevChar() rune {
	if i.position <= 0 {
		return 0
	}
	return i.runes[i.position-1]
}

func (i *LookaheadIterator) CurrChar() rune {
	return i.runes[i.position]
}

func (i *LookaheadIterator) NextChar() rune {
	if i.position >= len(i.runes)-1 {
		return 0
	}
	return i.runes[i.position+1]
}

func (i *LookaheadIterator) Chars() (prev rune, curr rune, next rune) {
	return i.PrevChar(), i.CurrChar(), i.NextChar()
}
