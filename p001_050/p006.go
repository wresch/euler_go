package main

import (
	eu "github.com/wresch/euler_go"
)

var desc string = `
The sum of the squares of the first ten natural numbers is,
\begin{equation}
	1^2 + 2^2 + ... + 10^2 = 385
\end{equation}
The square of the sum of the first ten natural numbers is,
\begin{equation}
	(1 + 2 + ... + 10)^22 = 55^2 = 3025
\end{equation}
Hence the difference between the sum of the squares of the first ten
natural numbers and the square of the sum is 3025 − 385 = 2640.

\emph{Find the difference between the sum of the squares of the first one
hundred natural numbers and the square of the sum.}

To simplify the sum of integers from 1 to $k$ if $k$ is even, we can rearrange 
\begin{equation}\begin{align}
\sum_{i=1}^k i &=& 1 + 2 + ... + (k - 1) + k \\
 &=& (1 + k) + (2 + (k - 1)) + ... \\
 &=& (1 + k) \frac{1}{2} k \\
 &=& \frac{(1 + k)k}{2}
\end{align}\end{equation}

Odd numbers can be partitioned in the sum up to the preceeding odd
number plus the number
\begin{equation}\begin{align}
\sum_{i=1}^k i &=& \sum_{i=1}^{k-1}i + k \\
 &=& \frac{(1 + k - 1) * (k - 1)}{2} + frac{2k}{2} \\
 &=& \frac{(1 + k)k}{2}
\end{align}\end{equation}

A more general proof for this solution is based on the
simple equality
\begin{equation}\begin{align}
(n + 1)^2 &=& n^2 + 2n + 1
(n + 1)^2 - n^2 &=& 2n + 1
\end{align}\end{equation}

We can write this out for all numbers from 1 to $k$
\begin{equation}\begin{align}
(1 + 1)^2 - 1^2 &=& 2*1 + 1 \\
(2 + 1)^2 - 2^2 &=& 2*2 + 1 \\
(3 + 1)^2 - 3^2 &=& 2*3 + 1 \\
... \\
(k + 1)^2 - k^2 &=& 2*k + 1
\end{align}\end{equation}

If we add up the terms on the left side, 
most terms cancel out except for $(k + 1)^2$ and
$-1$, which gives
\begin{equation}\begin{align}
(k + 1)^2 - 1 &=& 2\sum_{i=1}^k i + k \\
\sum_{i=1}^k i &=& \frac{(k + 1)^2 - 1 - k}{2} \\
\sum_{i=1}^k i &=& \frac{k(k + 1)}{2}
\end{align}\end{equation}

the same result as above.  The advantage of this
proof is that it can also be applied to other series. For example
\begin{equation}\begin{align}
(n + 1)^3 &=& n^3 + 3n^2 + 3n + 1 \\
(n + 1)^3 - n^3 &=& 3n^2 + 3n + 1 \\

(1 + 1)^3 - 1^3 &=& 3 * 1^2 + 3*1 + 1 \\
(2 + 1)^3 - 2^3 &=& 3 * 2^2 + 3*2 + 1 \\
... \\
(k + 1)^3 + k^3 &=& 3 * k^2 + 3*k + 1
\end{align}\end{equation}
Summing up left side and right side as before
\begin{equation}\begin{align}
(k + 1)^3 - 1 &=& 3\sum_{i=1}^k i^2 + 2\sum_{i=1}^k i + k
\sum_{i=1}^k i^2 &=& \frac{k(k + 1)(2k + 1)}{6}
\end{align}\end{equation}
`

func main() {
	eu.Output(6, desc, eu.Pow(eu.SumOfNaturalsUpTo(100), 2)-
		eu.SumOfNaturalSquaresUpTo(100))
}
