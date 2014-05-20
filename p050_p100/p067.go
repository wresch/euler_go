package main

import (
	"bufio"
	"fmt"
	eu "github.com/wresch/euler_go"
	"os"
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

\emph{Find the maximum total from top to bottom in p067.data.txt, a
15K text file containing a triangle with one-hundred rows.}

NOTE: This is a much more difficult version of Problem 18. It is not
possible to try every route to solve this problem, as there are $2^99$
altogether! If you could check one trillion ($10^12$) routes every second
it would take over twenty billion years to check them all. There is an
efficient algorithm to solve it.

My algorithm from problem 18 worked. Only loading data had to be changed.
`

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("p067.data.txt")
	check(err)
	defer file.Close()

	triangle := make([][]int, 0, 1000)
	scanner := bufio.NewScanner(file)
	ri := 0
	for scanner.Scan() {
		fields := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		fields_i := make([]int, ri+1)
		for ci, n := range fields {
			fields_i[ci], err = strconv.Atoi(n)
			check(err)
		}
		triangle = append(triangle, fields_i)
		ri++
	}
	check(scanner.Err())
	eu.Output(18, desc, fmt.Sprintf("Maximal sum of path through triangle is %d.\n",
		eu.TriangleMaxPathSum(triangle)))
}
