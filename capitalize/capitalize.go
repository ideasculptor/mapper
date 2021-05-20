package capitalize

import (
	"strings"
	"unicode"
)

func CapitalizeEveryThirdAlphanumericChar(s string) string {
	b := strings.Builder{}
	// Grow takes bytes, not a count of runes, so this is correct.
	// It's possible it needs 1 byte more, to hold a \0 terminator
	// without Growing again but I'd hope a builder is smart enough
	// not to need a terminator except when rendering the final
	// immutable string.
	b.Grow(len(s))

	alnumCount := 0
	// don't use index of range of string, which will provide
	// index into byte array instead of index into slice of runes
	for _, r := range strings.ToLower(s) {
		if (r >= '0' && r <= '9') || (r >= 'a' && r <= 'z') {
			alnumCount++
			if alnumCount%3 == 0 {
				r = unicode.ToUpper(r)
			}
		}
		b.WriteRune(r)
	}
	return b.String()
}
