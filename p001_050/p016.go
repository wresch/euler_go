package main

import (
	"fmt"
	"math/big"
	eu "github.com/wresch/euler_go"
)




var desc string = `

$2^15 = 32768$ and the sum of its digits is $3 + 2 + 7 + 6 + 8 = 26$.

\emph{What is the sum of the digits of the number $2^1000$?}

Instead of actually calculating the power, I created a big.Int,
set the 1000th bit, and the did the sum.
`

func main() {
	pow := big.NewInt(0)
	zero := big.NewInt(0)
	ten := big.NewInt(10)
	res := big.NewInt(0)
	mod := big.NewInt(0)
	pow.SetBit(pow, 1000, 1)
	for ;pow.Cmp(zero) == 1; {
		pow.DivMod(pow, ten, mod)
		res.Add(res, mod)
	}
	eu.Output(16, desc, fmt.Sprintf("Sum of digits of $2^1000$ is %v", res))
}