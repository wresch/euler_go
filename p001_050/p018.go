package main

import (
	"fmt"
	eu "github.com/wresch/euler_go"
	"strconv"
	"strings"
)

var desc string = `

By starting at the top of the triangle below and moving to adjacent
numbers on the row below, the maximum total from top to bottom is 23.

\begin{verbatim}
3
7 4
2 4 6
8 5 9 3
\end{verbatim}

That is, $3 + 7 + 4 + 9 = 23$.

\emph{Find the maximum total from top to bottom of the triangle below:}

\begin{verbatim}
75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23
\end{verbatim}

NOTE: As there are only 16384 routes, it is possible to solve this
problem by trying every route. However, Problem 67, is the same
challenge with a triangle containing one-hundred rows; it cannot be
solved by brute force, and requires a clever method!


Clearly, this cannot be solved by something as simple as always
taking the next step with the largest possible value. That goes wrong
with something as simple as
\begin{verbatim}
3
3 1
5 1 9
1 2 2 9
\end{verbatim}
where the incorrect answer would be $3 + 3 + 5 + 2 = 13$ instead of
the correct $3 + 1 + 9 + 9 = 22$.

I solved this by creating a corresponding tree in which each leaf
contained the maximal sum that can be obtained by moving down the
tree to that leaf. For the example above this would be
\begin{verbatim}
3
6   4
11  7 13
12 13 15 22 
\end{verbatim}

`

var sdata string = `
75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23
`

func main() {
	lines := strings.Split(strings.TrimSpace(sdata), "\n")
	input_n := make([][]int, len(lines))
	var err error
	for r, s := range lines {
		fields := strings.Split(strings.TrimSpace(s), " ")
		input_n[r] = make([]int, len(fields))
		for c, ns := range fields {
			input_n[r][c], err = strconv.Atoi(ns)
			if err != nil {
				panic(err)
			}
		}
	}
	eu.Output(18, desc, fmt.Sprintf("Maximal sum of path through triangle is %d.\n",
		eu.TriangleMaxPathSum(input_n)))
}
