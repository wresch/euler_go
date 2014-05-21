package euler_go

import (
	"fmt"
)

type PrimeSieve struct {
	Value      int64
	sieves     []*LinearCounter
	oddNumbers *LinearCounter
}

func NewPrimeSieve() (p *PrimeSieve) {
	p = &PrimeSieve{
		Value:      2,
		sieves:     make([]*LinearCounter, 0, 1000),
		oddNumbers: NewLinearCounter(3, 2),
	}
	return
}

func (s *PrimeSieve) Sieves() {
	for i, filt := range s.sieves {
		fmt.Printf("  sieve %3d: %6d\n", i, filt.Value)
	}
}

func (s *PrimeSieve) Next() int64 {
	for ; ; s.oddNumbers.Next() {
		isPrime := true
		val := s.oddNumbers.Value
		for _, filt := range s.sieves {
			if val == filt.Value {
				isPrime = false
				filt.Next()
			}
		}
		if isPrime {
			// new sieve advances by steps of two since even numbers are not
			// checked anyway
			s.Value = val
			s.sieves = append(s.sieves, NewLinearCounter(val*val, 2*val))
			s.oddNumbers.Next()
			return s.Value
		}
	}
}

func PrimeFactors(n int64) []int64 {
	prime_factors := make([]int64, 0, 100)
	for primes := NewPrimeSieve(); primes.Value*primes.Value <= n; primes.Next() {
		for n%primes.Value == 0 {
			prime_factors = append(prime_factors, primes.Value)
			n /= primes.Value
		}
	}
	if n != 1 {
		prime_factors = append(prime_factors, n)
	}
	return prime_factors
}

func SumOfDivisors(n int64) int64 {
	pfs := PrimeFactors(n)
	pfm := make(map[int64]int64)
	for _, pf := range pfs {
		pfm[pf] = pfm[pf] + 1
	}
	var sod int64 = 1
	for p, a := range pfm {
		sod *= (Pow(p, (a+1)) - 1) / (p - 1)
	}
	return sod
}
