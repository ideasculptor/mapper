package main

import (
	"fmt"
	"unicode"

	"mapper"
)

type SkipString struct {
	every      uint16
	runes      []rune
	countAtPos []uint16
}

func NewSkipString(every uint16, s string) SkipString {
	if every == 0 {
		// every is uint16 to avoid receiving a value too large for
		// the int index returned by range
		panic("Error: Invalid value for 'every': 0")
	}

	runes := []rune(s)
	ss := SkipString{
		every: every,
		runes: runes,
		// need len of []rune to avoid getting byte count of string
		// in case of multibyte codepoints
		countAtPos: make([]uint16, len(runes)),
	}

	// build a map of which indices should by uppercased so we don't
	// have to iterate from 0 to pos on every transformation
	alNumCount := uint16(0)
	for idx, r := range ss.runes {
		l := unicode.ToLower(r)
		switch {
		case (l >= '0' && l <= '9') || (l >= 'a' && l <= 'z'):
			alNumCount++
			ss.countAtPos[idx] = alNumCount
			if alNumCount == every {
				alNumCount = 0
			}
		default:
			// differentiate between alphanumerics and other codepoints
			// in case we only want to lower case the alphanumerics in future
			ss.countAtPos[idx] = 0
		}
	}
	return ss
}

func (s *SkipString) TransformRune(pos int) {
	if pos < 0 || pos >= len(s.runes) {
		panic("Error: pos out of range")
	}
	if s.countAtPos[pos] == s.every {
		s.runes[pos] = unicode.ToUpper(s.runes[pos])
	} else {
		// we could differentiate between alphanumerics that aren't 3rd
		// and other runes by checking if s.countAtPos[pos] == 0 here
		s.runes[pos] = unicode.ToLower(s.runes[pos])
	}
}

func (s *SkipString) GetValueAsRuneSlice() []rune {
	// defensively copy our slice of runes so that the caller
	// cannot mess with our internal state by modifying a slice
	// that shares an underlying array with our internal slice
	slice := make([]rune, len(s.runes))
	copy(slice, s.runes)
	return slice
}

func (s SkipString) String() string {
	return string(s.runes)
}

func main() {
	s := NewSkipString(3, "Aspiration.com")
	mapper.MapString(&s)
	fmt.Println(s)
}
