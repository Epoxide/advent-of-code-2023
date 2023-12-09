package decode

import (
	"testing"
)

func TestDecodeMultipleDigits(t *testing.T) {
  result := Decode("abc123")

	if result != 13 {
    t.Errorf("Decode(\"abc123\") = %d; want 13", result)
  }
}

func TestDecodeOnlyDigits(t *testing.T) {
  result := Decode("123")

	if result != 13 {
    t.Errorf("Decode(\"123\") = %d; want 13", result)
  }
}

func TestDecodeOneDigit(t *testing.T) {
  result := Decode("abc1")

	if result != 11 {
    t.Errorf("Decode(\"abc1\") = %d; want 11", result)
  }
}

func TestDecodeZero(t *testing.T) {
  result := Decode("abc00")

	if result != 0 {
    t.Errorf("Decode(\"abc00\") = %d; want 0", result)
  }
}

func TestDecodeLeadingZero(t *testing.T) {
  result := Decode("abc01")

	if result != 1 {
    t.Errorf("Decode(\"abc01\") = %d; want 1", result)
  }
}
