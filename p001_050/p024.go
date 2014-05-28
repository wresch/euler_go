package main

import (
	"fmt"
	eu "github.com/wresch/euler_go"
)

var desc = `
A permutation is an ordered arrangement of objects. For example, 3124
is one possible permutation of the digits 1, 2, 3 and 4. If all of the
permutations are listed numerically or alphabetically, we call it
lexicographic order. The lexicographic permutations of 0, 1 and 2 are:

\begin{verbatim}
012   021   102   120   201   210
\end{verbatim}

\emph{What is the millionth lexicographic permutation of the digits 0, 1, 2,
3, 4, 5, 6, 7, 8 and 9?}

\subsection{Approach}

There are $10!=3,628,800$ permutations of the digits 0-9 (since order
matters).  The smallest permutation is $0,123,456,789$. The $9! =
362,880$th permutation is the largest permutation that still starts
with a 0, so that would be $0,987,654,321$. Permutation $2\cdot9! =
725760$ is $1,987,654,320$ and clearly the millionth permutation would
have to be larger, so it starts with the digit 2. Repeating the same
reasoning for each successive step up to permutation $2\cdot9! +
6\cdot8! + 6\cdot7! + 2\cdot6! + 5\cdot5! + 1\cdot4! + 2\cdot3! =
999996$ is $2,783,914,650$.  Four more steps take us to
$2,783,915,460$.  `


func main() {
	eu.Output(24, desc, fmt.Sprintf("The millionth permutation of 0-9 is %d.\n", 2783915460))
}
