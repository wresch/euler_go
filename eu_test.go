package euler_go

import (
	"fmt"
	"testing"
)

var british_number_tests = []struct {
	n int64
	s string
}{
	{1, "one"},
	{11, "eleven"},
	{17, "seventeen"},
	{21, "twenty-one"},
	{-33, "negative thirty-three"},
	{96, "ninety-six"},
	{101, "one hundred and one"},
	{120, "one hundred and twenty"},
	{123, "one hundred and twenty-three"},
	{548, "five hundred and forty-eight"},
	{1367, "one thousand three hundred and sixty-seven"},
	{5409, "five thousand four hundred and nine"},
	{312003, "three hundred and twelve thousand and three"},
	{400000, "four hundred thousand"},
	{532157, "five hundred and thirty-two thousand one hundred and fifty-seven"},
	{1000001, "one million and one"},
}

func TestBritishSpelledNumber(t *testing.T) {
	for _, test := range british_number_tests {
		s := BritishSpelledNumber(test.n)
		if s != test.s {
			t.Error(fmt.Sprintf("For '%d' expected '%s' but got '%s'",
				test.n, test.s, s))
		}
	}
}
