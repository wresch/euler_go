package main

import (
	"fmt"
	eu "github.com/wresch/euler_go"
	"math/big"
)

var desc string = `


$n!$ means $n \cdot (n − 1) \cdot ... \cdot 3 \cdot 2 \cdot 1$

For example, $10! = 10 \cdot 9 \cdot ... \cdot 3 \cdot 2 cdot 1 = 3628800$,
and the sum of the digits in the number $10!$ is $3 + 6 + 2 + 8 + 8 + 0 + 0 = 27$.

\emph{Find the sum of the digits in the number 100!}

I'm not aware of a clever solution, so I just use big numbers!
`

func main() {
	N := 100
	var sumOfDigits int64 = 0
	fact100 := big.NewInt(0)
	fact100.MulRange(int64(1), int64(N))
	fact100s := fact100.String()
	mod10 := big.NewInt(0)
	ten := big.NewInt(10)
	zero := big.NewInt(0)
	for fact100.Cmp(zero) == 1 {
		fact100.DivMod(fact100, ten, mod10)
		sumOfDigits += mod10.Int64()
	}
	eu.Output(20, desc, fmt.Sprintf("Sum of digits of $100!$ is %d\nNumber:",
		sumOfDigits))
	fmt.Println("\\begin{verbatim}")
	i := 0
	for ; i+50 <= len(fact100s); i += 50 {
		fmt.Println(fact100s[i:(i + 50)])
	}
	if i < len(fact100s) {
		fmt.Println(fact100s[i:])
	}
	fmt.Println("\\end{verbatim}")
}
