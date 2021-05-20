package capitalize

import (
	"fmt"
	"strings"
	"unicode"
)

func CapitalizeEveryThirdAlphanumericChar(s string) string {
	// I'd probably use []rune here, but I included the strings.Builder
	// just to demonstrate I know it exists
	output := make([]rune, 0, len(s))
	b := strings.Builder{}
	b.Grow(len(s))

	alnumCount := 0
	// don't use index of range of string, which will provide
	// index into byte array instead of slice of runes
	for _, r := range strings.ToLower(s) {
		if (r >= '0' && r <= '9') || (r >= 'a' && r <= 'z') {
			alnumCount++
			if alnumCount%3 == 0 {
				r = unicode.ToUpper(r)
			}
		}
		output = append(output, r)
		b.WriteRune(r)
	}
	fmt.Printf("output: %s\n", string(output))
	fmt.Printf("builder: %s\n", b.String())
	return b.String()
}
