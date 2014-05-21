package main

import (
	"fmt"
	eu "github.com/wresch/euler_go"
)

var desc string = `


Let $d(n)$ be defined as the sum of proper divisors of n (numbers less
than n which divide evenly into n).  If $d(a) = b$ and $d(b) = a$, where
$a \neq b$, then $a$ and $b$ are an amicable pair and each of $a$ and $b$
are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20,
22, 44, 55 and 110; therefore $d(220) = 284$. The proper divisors of 284
are 1, 2, 4, 71 and 142; so $d(284) = 220$.

\emph{Evaluate the sum of all the amicable numbers under 10000.}

\subsection{Approach}
$d(n)$ is the sum-of-divisors function also called $\sigma_{1}(n)$. This is
a special case of a divisor function $\sigma_{x}(n)$ which is the sum of the xth
powers of all divisors of a number $n$. $\sigma_{0}(n)$ is the number-of-divisors
function discussed in problem 12.  For a prime power $n = p^{a}$, $\sigma_{1}(n)$
can be calculated as follows:
\[
\sigma_{1}(n) = 1 + p + p^{2} + ... + p^{a} \tag{I}
\]
Multiply both side by $p$
\[
p\sigma_{1}(n) = p + p^{2} + ... + p^{a+1} \tag{II}
\]
and subtract (I) from (II)
\begin{align*}
\sigma_{1}(n)\cdot(p - 1) = p^{a + 1} - 1 \\
\sigma_{1}(n) = \frac{p^{a + 1} - 1}{p - 1} \tag{III}
\end{align*}

Now, take two prime powers $n_{1} = p^{a}$ and $n_{2} = q^{b}$ where $p \neq q$.
Based on (I):
\begin{align*}
\sigma_{1}(n_{1}) = 1 + p + p^{2} + ... + p^{a} \\
\sigma_{1}(n_{2}) = 1 + q + q^{2} + ... + q^{b}
\end{align*}

For any given $q^{i}$
\[
q^{i}(1 + p + p^{2} + ... + p^{a}) = q^{i} \sigma_{1}(n_{1})
\]

Summed over all $q^{i}$

\begin{align*}
(1 + q^{1} + ... q^{b})(1 + p + ... + p^{a}) = (1 + q^{1} + ... q^{b}) \sigma_{1}(n_{1}) \\
(1 + q^{1} + ... q^{b})(1 + p + ... + p^{a}) = \sigma_{1}(n_{1})\sigma_{1}(n_{2})
\end{align*}

We can see that that the left side equals
\[
q^{0}p^{0} + q^{0}p^{1} + .. + q^{0}p^{a} + q^{1}p^{0} + q^{1}p^{1} + ...
\]

which is the sum of all divisors of the number $n = n_{1}n_{2}$ and therefore
\[
\sigma_{1}(n_{1}n_{2}) = \sigma_{1}(n_{1})\sigma_{1}(n_{2})
\]

and more generally, for a number $N$ factored into prime powers
$N=p_{1}^{a_{1}}p_{2}^{a_{2}}...p_{n}^{a_{n}}$
\[
\sigma_{1}(N) = \prod_{i=1}^{n}\frac{p_{i}^{a_{i} + 1} - 1}{p_{i} - 1}
\]
where $n$ is the number of prime divisors.
`

func main() {
	var N int64 = 10000
	sods := make([]int64, 10001)
	for i := int64(1); i <= N; i++ {
		sods[i] = eu.SumOfDivisors(i) - i // only using proper divisors
	}
	var sumOfAmicableNumbers int64 = 0
	for i, sod := range sods {
		i64 := int64(i)
		if i64 < N && sod < N && sod > 0 {
			if sods[sod] == i64 && i64 != sod {
				//fmt.Printf("%4d[%d] - %4d[%d]\n", i, sod, sod, sods[sod])
				sumOfAmicableNumbers += i64 + sod
				sods[i] = 0
				sods[sod] = 0
			}
		}
	}
	eu.Output(21, desc,
		fmt.Sprintf("The sum of all amicable numbers below 10000 is %d.\n",
			sumOfAmicableNumbers))
}
