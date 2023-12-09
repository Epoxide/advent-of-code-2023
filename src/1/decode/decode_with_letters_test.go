package decode

import (
	"testing"
)

func TestDecodeWithLettersMultipleDigits(t *testing.T) {
  result := DecodeWithLetters("abc123")

	if result != 13 {
    t.Errorf("DecodeWithLetters(\"abc123\") = %d; want 13", result)
  }
}

func TestDecodeWithLettersOnlyDigits(t *testing.T) {
  result := DecodeWithLetters("123")

	if result != 13 {
    t.Errorf("DecodeWithLetters(\"123\") = %d; want 13", result)
  }
}

func TestDecodeWithLettersOneDigit(t *testing.T) {
  result := DecodeWithLetters("abc1")

	if result != 11 {
    t.Errorf("DecodeWithLetters(\"abc1\") = %d; want 11", result)
  }
}

func TestDecodeWithLettersZero(t *testing.T) {
  result := DecodeWithLetters("abc00")

	if result != 0 {
    t.Errorf("DecodeWithLetters(\"abc00\") = %d; want 0", result)
  }
}

func TestDecodeWithLettersLeadingZero(t *testing.T) {
  result := DecodeWithLetters("abc01")

	if result != 1 {
    t.Errorf("DecodeWithLetters(\"abc01\") = %d; want 1", result)
  }
}

func TestDecodeWithLettersWords(t *testing.T) {
  result := DecodeWithLetters("onextwoxthreex1xnine")

	if result != 19 {
    t.Errorf("DecodeWithLetters(\"onextwoxthreex1xnine\") = %d; want 19", result)
  }

	result = DecodeWithLetters("two1nine")

	if result != 29 {
    t.Errorf("DecodeWithLetters(\"two1nine\") = %d; want 29", result)
  }

	result = DecodeWithLetters("eightwothree")

	if result != 83 {
    t.Errorf("DecodeWithLetters(\"eightwothree\") = %d; want 83", result)
  }

	result = DecodeWithLetters("abcone2threexyz")

	if result != 13 {
    t.Errorf("DecodeWithLetters(\"abcone2threexyz\") = %d; want 13", result)
  }

	result = DecodeWithLetters("xtwone3four")

	if result != 24 {
    t.Errorf("DecodeWithLetters(\"xtwone3four\") = %d; want 24", result)
  }

	result = DecodeWithLetters("4nineeightseven2")

	if result != 42 {
    t.Errorf("DecodeWithLetters(\"4nineeightseven2\") = %d; want 42", result)
  }

	result = DecodeWithLetters("zoneight234")

	if result != 14 {
    t.Errorf("DecodeWithLetters(\"zoneight234\") = %d; want 14", result)
  }

	result = DecodeWithLetters("7pqrstsixteen")

	if result != 76 {
    t.Errorf("DecodeWithLetters(\"7pqrstsixteen\") = %d; want 76", result)
  }

	result = DecodeWithLetters("qnqeightwofcdsxgscgclpptnp9four1sixnjlvxqxxsnine")

	if result != 89 {
    t.Errorf("DecodeWithLetters(\"qnqeightwofcdsxgscgclpptnp9four1sixnjlvxqxxsnine\") = %d; want 89", result)
  }

	result = DecodeWithLetters("two")

	if result != 22 {
    t.Errorf("DecodeWithLetters(\"two\") = %d; want 22", result)
  }
}

func TestDecodeWithLettersWordsEndingInCombination(t *testing.T) {
	result := DecodeWithLetters("onetwooneight")

	if result != 18 {
    t.Errorf("DecodeWithLetters(\"onetwooneight\") = %d; want 18", result)
  }
}
