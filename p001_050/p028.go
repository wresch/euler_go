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

`

func main() {
	eu.Output(28, desc, fmt.Sprintf(`
The sum of the diagonals of a 1001 by 1001 spiral is %d.\n`, 0))
}
