package main

import (
	"fmt"
	eu "github.com/wresch/euler_go"
)

var desc string = `
If the numbers 1 to 5 are written out in words: one, two, three, four,
five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

\emph{If all the numbers from 1 to 1000 (one thousand) inclusive were
written out in words, how many letters would be used?}

NOTE: Do not count spaces or hyphens. For example, 342 (three hundred
and forty-two) contains 23 letters and 115 (one hundred and fifteen)
contains 20 letters. The use of "and" when writing out numbers is in
compliance with British usage.

This was implemented with a dictionary based number generation
recursive function followed by actual counting of letters. It is
possible to solve this on paper by observing some simple patterns, but
the function is interesting to implement anyway.  `

func main() {
	total := 0
	for i := int64(1); i <= 1000; i++ {
		s := eu.BritishSpelledNumber(i)
		sl := 0
		for _, l := range s {
			if l != ' ' && l != '-' {
				sl++
			}
		}
		total += sl
	}
	eu.Output(17, desc, fmt.Sprintf("%d letters were used to spell numbers 1 - 1000", total))
}
