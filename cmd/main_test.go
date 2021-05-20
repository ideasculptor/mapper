package main

import (
	"mapper"
	"testing"
)

func TestPrimaryCase(t *testing.T) {
	input := "Aspiration.com"
	expected := "asPirAtiOn.cOm"
	s := NewSkipString(3, input)
	mapper.MapString(&s)

	if s.String() != expected {
		t.Errorf("NewSkipString(3, %s) = %s, expected: %s", input, s, expected)
	}
}

func TestDifferentEvery(t *testing.T) {
	input := "Aspiration.com"
	expected := "aspIratIon.cOm"

	s := NewSkipString(4, input)
	mapper.MapString(&s)

	if s.String() != expected {
		t.Errorf("NewSkipString(4, %s) = %s, expected: %s", input, s, expected)
	}
}

func TestMultiByteWithoutAlNum(t *testing.T) {
	input := "Очень хороший"
	expected := "очень хороший"

	s := NewSkipString(3, input)
	mapper.MapString(&s)

	if s.String() != expected {
		t.Errorf("NewSkipString(3, %s) = %s, expected: %s", input, s, expected)
	}
}

func TestMultiByteWithAlNum(t *testing.T) {
	input := "ОчaBень х1оcрDоfший"
	expected := "очabень х1оcрdоFший"

	s := NewSkipString(3, input)
	mapper.MapString(&s)

	if s.String() != expected {
		t.Errorf("NewSkipString(3, %s) = %s, expected: %s", input, s, expected)
	}
}

func TestMultiByteWithAlNumDifferentEvery(t *testing.T) {
	input := "Очabень х1оcрDоfший"
	expected := "очaBень х1оCрdоFший"

	s := NewSkipString(2, input)
	mapper.MapString(&s)

	if s.String() != expected {
		t.Errorf("NewSkipString(2, %s) = %s, expected: %s", input, s, expected)
	}
}
