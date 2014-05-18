package main

import (
	"fmt"
	"math/big"
	eu "github.com/wresch/euler_go"
)

var desc string = `

Starting in the top left corner of a 2x2 grid, and only being able to
move to the right and down, there are exactly 6 routes to the bottom
right corner.

\emph{How many such routes are there through a 20x20 grid?}

There are multiple solutions for this.  Two solution I implemented
here are an iterative solution that saves intermediate steps and a
combinatorial solution.

\subsection{Iterative solution}

Starting on a grid at position $(0, 0)$, there are two initial steps
towards the right bottom corner $(m, n)$ one can take: Towards $(1, 0)$ and
towards $(0, 1)$. Therefore
\[
N(0, 0) = N(1, 0) + N(0, 1)
\]

To avoid recalculating solutions, intermediate solutions are
saved in a 2 dimensional table, where all $(m, 0)$ and $(0, n)$ values
are initialized to 1 and the calculations start at $(m - 1, n - 1)$.

\subsection{Combinatorial solution}

Each march from $(0, 0)$ to $(m, n)$ involves exactly m steps down
and n steps to the right for a total of $m + n$ steps.  Each march
can be described as a random strings of $R$ight and $D$own steps.
Once we place all the $R$s, we know where all the $D$s go. So, the
number of possible routes is
\[
{n + m \choose n} = \frac{(n + m)!}{n!\cdot(n + m - n)!} = \frac{(n + m)!}{n!\cdot m!} \\
{n + m \choose m} = \frac{(n + m)!}{m!\cdot(n + m - m)!} = \frac{(n + m)!}{m!\cdot n!} \\
{n + m \choose n} = {n + m \choose m} = \frac{(m + n)\cdot(m + n - 1)\cdot...\cdot(n + 1)}{m!}
\]


 `
// a 2x2 grid has 3 points in each row!
func path_numbers_iterative(m, n int) int64 {
	grid := make([][]int64, m + 1)
	for r := 0; r <= m; r++ {
		grid[r] = make([]int64, n + 1)
	}
	for r := 0; r <= m; r++ {
		grid[r][n] = 1
	}
	for c := 0; c <= n; c++ {
		grid[m][c] = 1
	}
	for r := m - 1; r >= 0; r-- {
		for c := n - 1; c >= 0; c-- {
			grid[r][c] = grid[r+1][c] + grid[r][c+1]
		}
	}
	//for r := 0; r < m; r++ {
	//	for c:= 0; c < n; c++ {
	//		fmt.Printf("%12d", grid[c][r])
	//	}
	//	fmt.Println("")
	//}
	return grid[0][0]
}

func path_numbers_combinatorial(m, n int) *big.Int {
	m64 := int64(m)
	n64 := int64(n)
	r := big.NewRat(m64 + n64, m64)
	for k := int64(1); k < m64; k++ {
		r.Mul(r, big.NewRat(n64 + m64 - k, m64 - k))
	}
	if r.IsInt() {
		//fmt.Printf("r.Num() -> %p\n", r.Num())
		r_int := *r.Num()
		//fmt.Printf("r_int   -> %p\n", &r_int)
		return &r_int
	} else {
		panic("Number of paths was not an integer")
	}
}

func main() {
	r1 := path_numbers_iterative(20, 20)
	r2 := path_numbers_combinatorial(20, 20)
	eu.Output(15, desc, fmt.Sprintf("\\begin{tabular}{rr}\nIterative solution       & %25d\\\\\nCombinatorial solution   & %25s\n\\end{tabular}\n", r1, r2))
}