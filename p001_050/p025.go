package main

import (
	"fmt"
	"math/big"
	eu "github.com/wresch/euler_go"
)

var desc = `

The Fibonacci sequence is defined by the recurrence relation:
\[
    F_n = F_{n−1} + F_{n−2}, where F_1 = 1 and F_2 = 1
\]

Hence the first 12 terms will be:
\begin{align*}
    F1 &= 1 \\
    F2 &= 1 \\
    ...
    F12 = 144 \\
\end{align*}

The 12th term, $F_12$, is the first term to contain three digits.

\emph{What is the first term in the Fibonacci sequence to contain 1000
digits?}

\subsection{Approach}

It can be shown that the ratio of consecutive Fibonacci numbers
approaches the golden ratio $\phi$:

\begin{align*}
\lim_{n \to \infty}\frac{F_n}{F_{n-1}} &= \lim_{n \to \infty}\frac{F_{n-1}+F_{n-2}}{F_{n-1}} \\
&= \lim_{n \to \infty}(1 + \frac{F_{n-2}}{F_{n-1}})
\end{align*}

with $X = \frac{F_n}{F_{n-1}}$ for large $n$ we get
\begin{align*}
X &= 1 + \frac{1}{X} \\
X^2 - X - 1 &= 0 \\
X = \frac{1+\sqrt{5}}{2}
\end{align*}
which equals the golden ratio $\phi = 1.61803...$.

Now, the smallest number with 1000 digits is $10^999$.  If we approximate the $n$th
Fibonacci number starting with the 10th number (55) as a power of $\phi$ we can say that
\[
10^{999} < 55\cdot\phi^{n-10}
\]

which we can solve
\begin{array*}
999 &< log_{10}55 + (n - 10)log_{10}\phi \\
999 - log_{10}55 &< n\log_{10}\phi - 10\log_{10}\phi \\
n\log_{10}\phi &> 999 - log_{10}55 + 10\log_{10}\phi \\
n &> \frac{999 - log_{10}55 + 10\log_{10}\phi}{\log_{10}\phi}
\end{array*}

resulting in $n > 4781.9$, i.e. $n = 4782$.

Since this is a programming challenge, i also brute forced it with golangs
arbitrary precision package, which is a bit cumersome.

\subsection{Quadratic equation proof reminder}
\begin{align*}
ax^2+bx+c &= 0 \\
x^2 + \frac{b}{a}x + \frac{c}{a} &= 0 \\
x^2 + \frac{b}{a}x &= -\frac{c}{a} \\
x^2 + \frac{b}{a}x + (\frac{b}{2a})^2 &= (\frac{b}{2a})^2 -\frac{c}{a} \\
(x + \frac{b}{2a})^2 &= \frac{b^2-4ca}{4a^2} \\
x &= -\frac{b}{2a}\pm \sqrt\frac{b^2-4ca}{4a^2} \\
x &= -\frac{b}{2a}\pm \frac{\sqrt{b^2-4ac}}{2a} \\
x &= \frac{-b\pm\sqrt{b^2-4ac}}{2a}
\end{align*}
`




func main() {
	eu.Output(25, desc, fmt.Sprintf("The first Fibonacci number with 1000 digits is the %dth term of the series (mathematically derived).", 4782))
	kdigits := big.NewInt(0)
	kdigits.Exp(big.NewInt(10), big.NewInt(999), nil)
	fibTerms := [3]*big.Int{big.NewInt(1), big.NewInt(1), big.NewInt(2)}
	n := 3
	for fibTerms[2].Cmp(kdigits) == -1 {
		n++
		fibTerms[0].Set(fibTerms[1])
		fibTerms[1].Set(fibTerms[2])
		fibTerms[2].Add(fibTerms[0], fibTerms[1])
	}
	fmt.Printf("By brute force we also find %d.\n", n)	
}
