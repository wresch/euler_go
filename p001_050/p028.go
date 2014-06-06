package main

import (
	"fmt"
	eu "github.com/wresch/euler_go"
)

var desc = `


Starting with the number 1 and moving to the right in a clockwise
direction a 5 by 5 spiral is formed as follows:

\begin{verbatim}
21 22 23 24 25
20  7  8  9 10
19  6  1  2 11
18  5  4  3 12
17 16 15 14 13
\end{verbatim}

It can be verified that the sum of the numbers on the diagonals is
101.

\emph{What is the sum of the numbers on the diagonals in a 1001 by
1001 spiral formed in the same way?}

\subsection{Approach}

Basically for each additional turn of the spiral, we have to add
4 more numbers.  For example for the first turn of the spiral we
add
\[
3^2 + 3^2 - 2 + 3^2 - 2\cdot2 + 3^2 - 3\cdot2 = 4\cdot3^2 - 6\cdot(3-1)
\]

This can be generalized for spirals of width n
\begin{align*}
ds_n &= 1 + \sum_{i=0}^{(n-3)/2}4(2i+3)^2 - 6(2i+2)
ds_n &= 1 + \sum_{i=0}^{(n-3)/2}16i^2 + 36i + 24
\end{align*}

Given that
\[
\sum_{i=0}^{m}i = \frac{m(m+1)}{2}
\]
and
\[
\sum_{i=0}^{m}i^2 = \frac{m(m+1)(2m+1)}{6}
\]
it should be possible to come up with a closed form solution
to the summartion, but apparently my arithmetic is not up to the
task.  So instead, I just implemented the sum above

`

func sumOfDiagonals(squareSize int) (r int) {
	r = 1
	for i := 0; i <= (squareSize-3)/2; i++ {
		r += 16*i*i + 36*i + 24
	}
	return
}

func main() {
	eu.Output(28, desc, fmt.Sprintf(`
The sum of the diagonals of a 1001 by 1001 spiral is %d.`, sumOfDiagonals(1001)))
}
