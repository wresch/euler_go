package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	fact := big.NewInt(0).MulRange(2, int64(N))
	facts := fact.String()
	i := 0
	for ; i+50 <= len(facts); i += 50 {
		fmt.Println(facts[i:(i + 50)])
	}
	if i < len(facts) {
		fmt.Println(facts[i:])
	}
}
