package capitalize

import (
	"testing"
)

func TestPrimaryCase(t *testing.T) {
	input := "Aspiration.com"
	expected := "asPirAtiOn.cOm"

	s := CapitalizeEveryThirdAlphanumericChar(input)

	if s != expected {
		t.Errorf("CapitalizeEveryThirdAlphanumericChar(%s) = %s, expected: %s", input, s, expected)
	}
}

func TestMultiByteWithoutAlNum(t *testing.T) {
	input := "Очень хороший"
	expected := "очень хороший"

	s := CapitalizeEveryThirdAlphanumericChar(input)

	if s != expected {
		t.Errorf("CapitalizeEveryThirdAlphanumericChar(%s)) = %s, expected: %s", input, s, expected)
	}
}

func TestMultiByteWithAlNum(t *testing.T) {
	input := "ОчaBень х1оcрDоfший"
	expected := "очabень х1оcрdоFший"

	s := CapitalizeEveryThirdAlphanumericChar(input)

	if s != expected {
		t.Errorf("CapitalizeEveryThirdAlphanumericChar(%s) = %s, expected: %s", input, s, expected)
	}
}
