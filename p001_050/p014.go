package main

import (
	"fmt"
	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
	"image/color"
	eu "github.com/wresch/euler_go"
)

var desc string = `
The following iterative sequence is defined for the set of positive
integers:

\[
f(n) = \begin{cases} n/2 &\mbox{if } n \equiv 0 \\
(3n +1)/2 & \mbox{if } n \equiv 1. \end{cases} \pmod{2}
\]

Using the rule above and starting with 13, we generate the following
sequence:

\[
13 \to 40 \to 20 \to 10 \to 5 \to 16 \to 8 \to 4 \to 2 \to 1
\]

It can be seen that this sequence (starting at 13 and finishing at 1)
contains 10 terms. Although it has not been proved yet (Collatz
Problem), it is thought that all starting numbers finish at 1.

\emph{Which starting number, under one million, produces the longest
chain?}

NOTE: Once the chain starts the terms are allowed to go above one
million.

I don't think that there are any big optimizations or shortcuts
to be had other than starting from the top.  However, i'm making
this a bit more interesting by checking the length for all
numbers below $10^6$ and creating a histogram of their corresponding
chain
`

//////////////////// problem 14 ///////////////////

func main() {
	var N int64 = 1000000
	//collatz_lengths := make([]int64, N)
	collatz_lengths := make(plotter.Values, N)
	var max_collatz_length int64 = 0
	var max_collatz_int int64 = 0
	for i := int64(1); i <= N; i++ {
		cs := eu.NewCollatzSeries(i)
		for ; cs.Value != 1; {
			cs.Next()
		}
		collatz_lengths[i-1] = float64(cs.Length)
		if cs.Length > max_collatz_length {
			max_collatz_length = cs.Length
			max_collatz_int = i
		}
	}
	eu.Output(14, desc, fmt.Sprintf("Longest series starts at %d and has length %d\n",
		max_collatz_int, max_collatz_length))

	// create histogram
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Collatz series length for values 1-10^6"
	
	h, err := plotter.NewHist(collatz_lengths, 100)
	if err != nil {
		panic(err)
	}
	h.FillColor = color.RGBA{R: 190, G: 190, B: 190}
	p.Add(h)
	if err := p.Save(4, 4, "p014.pdf"); err != nil {
		panic(err)
	}
}
