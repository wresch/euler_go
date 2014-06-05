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

////////////////////////////////////////////////////////////////////////////////
// Data structure to store prime numbers efficiently
//   takes advantage of the fact that primes > 3 can be expressed as either
//   p = 6n + 1 or p = 6n + 5.  Proof for this is shown elsewhere. Originally
//   i wanted to use big.Int as a bitset, but it was too slow (too many
//   allocations). Instead i wrote a minimal bitset myself
////////////////////////////////////////////////////////////////////////////////
type PrimeSet struct {
	modulo6 []uint64
	max     uint
}

func NewPrimeSet(max uint) (p *PrimeSet) {
	max64 := int64(max)
	p = &PrimeSet{
		modulo6: make([]uint64, max/64+1),
		max:     max,
	}
	var i uint = 0
	for psieve := NewPrimeSieve(); psieve.Value < max64; psieve.Next() {
		i = 0
		if psieve.Value < 5 {
			continue
		}
		d6 := uint(psieve.Value) / 6
		m6 := uint(psieve.Value) % 6
		switch m6 {
		case 1:
			i = d6
		case 5:
			i = d6 + 1
		default:
			panic("Bad prime number")
		}
		p.modulo6[i>>6] |= 1 << (i & 63)
	}
	return p
}

func (p *PrimeSet) IsPrime(n uint) (isPrime, ok bool) {
	isPrime, ok = false, false
	if n < p.max {
		var i uint = 0
		switch n {
		case 2, 3, 5:
			isPrime, ok = true, true
		default:
			d6 := n / 6
			m6 := n % 6
			switch m6 {
			case 1:
				i = d6
			case 5:
				i = d6 + 1
			default:
				isPrime, ok = false, true
				return
			}
			if p.modulo6[i>>6]&(1<<(i&63)) != 0 {
				isPrime, ok = true, true
			}
		}
	}
	return
}

////////////////////////////////////////////////////////////////////////////////
// Prime factorization of a number n
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

////////////////////////////////////////////////////////////////////////////////
// Sum of all divisors (including 1 and n itself) of a number n
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
