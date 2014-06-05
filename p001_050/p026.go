package main

import (
	"bytes"
	"fmt"
	eu "github.com/wresch/euler_go"
)

var desc = `
A unit fraction contains 1 in the numerator. The decimal
representation of the unit fractions with denominators 2 to 10 are
given:

\begin{align*}
    \frac{1}{2} &= 0.5 \\
    \frac{1}{3} &= 0.\overline{3} \\
    \frac{1}{4} &= 0.25 \\
    \frac{1}{5} &= 0.2 \\
    \frac{1}{6} &= 0.1\overline{6} \\
    \frac{1}{7} &= 0.\overline{142857} \\
    ...
\end{align*}

Where $0.1\overline{6}$ has a 1-digit recurring cycle. It can be
seen that $\slfrac{1}{7}$ has a 6-digit recurring cycle.

Find the value of d < 1000 for which $\slfrac{1}{d}$ contains the longest
recurring cycle in its decimal fraction part.

\subsection{Approach}

First, which numbers don't have a recurring cycle? These should be all
numbers that evenly divide a power of 10, which in turn means that
they can be factored into the form
\[
d = 2^a + 5^b
\]
For these it won't be necessary to determine the length of the
cycle. For all other numbers, we implement a basic long division
to determine the period of the recurring cycle.  We start looking from
the largest possible number.

I think i am missing some interesting mathematics here, but it
works.
`

func unitFractionTerminates(d int) bool {
	pf := eu.PrimeFactors(int64(d))
	isTerminating := true
	for _, v := range pf {
		if v != 2 && v != 5 {
			isTerminating = false
			break
		}
	}
	return isTerminating
}

func unitFractionDigits(d int) (repeatStart int, digits []int) {
	digits = make([]int, 0, 100)
	remainders := make(map[int]int)
	n := 10
	repeatStart = -1
	for i := 0; ; i++ {
		rem := n % d
		dig := n / d
		if pos, ok := remainders[rem]; ok && digits[pos] == dig {
			repeatStart = pos
			break
		}
		remainders[rem] = i
		digits = append(digits, dig)
		if rem == 0 {
			break
		}
		n = rem * 10
	}
	return repeatStart, digits
}

func digitsAsString(digits []int, repeatStart int) string {
	var buf bytes.Buffer // buffers need no initialization
	if repeatStart == -1 {
		for _, v := range digits {
			buf.WriteByte(byte(v + 48))
		}
	} else {
		for i, v := range digits {
			if i == repeatStart {
				buf.WriteByte('(')
			}
			buf.WriteByte(byte(v + 48))
		}
		buf.WriteByte(')')
	}
	return buf.String()
}

func main() {
	maxRepeatCycle := 0
	maxRepeatCycleNum := 0
	for i := 999; i > 2; i++ {
		if !unitFractionTerminates(i) {
			repStart, digits := unitFractionDigits(i)
			if len(digits)-repStart > maxRepeatCycle {
				maxRepeatCycle = len(digits) - repStart
				maxRepeatCycleNum = i
			}
			//fmt.Printf("1/%-3d repeats:    0.%s (cycle: %d)\n", i, digitsAsString(digits, repStart), len(digits) - repStart)
		}
	}
	eu.Output(26, desc, fmt.Sprintf(`The number d < 1000 with the longest recurring cycle
in the decimal representation of its unit fraction is %d with a cycle length of %d.\n`, maxRepeatCycleNum, maxRepeatCycle))
}
