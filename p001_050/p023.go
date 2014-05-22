package main

import (
	"fmt"
	eu "github.com/wresch/euler_go"
	"sort"
)

var desc = `
A perfect number is a number for which the sum of its proper divisors
is exactly equal to the number. For example, the sum of the proper
divisors of 28 would be $1 + 2 + 4 + 7 + 14 = 28$, which means that 28
is a perfect number.

A number $n$ is called deficient if the sum of its proper divisors is
less than $n$ and it is called abundant if this sum exceeds $n$.

As 12 is the smallest abundant number, $1 + 2 + 3 + 4 + 6 = 16$, the
smallest number that can be written as the sum of two abundant numbers
is 24. By mathematical analysis, it can be shown that all integers
greater than 28123 can be written as the sum of two abundant
numbers. However, this upper limit cannot be reduced any further by
analysis even though it is known that the greatest number that cannot
be expressed as the sum of two abundant numbers is less than this
limit.

\emph{Find the sum of all the positive integers which cannot be
written as the sum of two abundant numbers.}

\subsection{Approach}

Using the previously derived approach for obtaining the sum of
proper divisors, all abundant numbers below 28123 - 12 are
identified. Since there are quite a few abundant numbers it
should be more efficient to go through each number below
28123 and find if it can be written as a sum of abundant
numbers rather than forming all pairs of abundant numbers and
seeing which numbers are not part of the results.

One optimization is that for a number $i$, we only need to
check against abundant numbers below $\frac{i}{2}$.
`

func main() {
	abundantNums := make(map[int]int)
	i := 12
	for ; i < 28123; i++ {
		if int(eu.SumOfDivisors(int64(i)))-i > i {
			abundantNums[i] = 1
		}
	}
	//fmt.Printf("Found %d abundant numbers < 28123\n", len(abundantNums))
	abundantNumsList := make([]int, len(abundantNums))
	i = 0
	for k, _ := range abundantNums {
		abundantNumsList[i] = k
		i++
	}
	sort.Ints(abundantNumsList)
	notSumOf2ANs := make([]int, 23, 1000)
	i = 1
	for ; i < 24; i++ {
		notSumOf2ANs[i-1] = i
	}
	for ; i < 28123; i++ {
		isSum := false
		for _, n := range abundantNumsList {
			if n > i/2 {
				break
			}
			if abundantNums[i-n] == 1 {
				isSum = true
				break
			}
		}
		if !isSum {
			notSumOf2ANs = append(notSumOf2ANs, i)
		}
	}
	total := 0
	for _, v := range notSumOf2ANs {
		total += v
	}
	eu.Output(23, desc, fmt.Sprintf("Sum of all numbers that are not the sum of two abundant numbers is %d.\n",
		total))
}
