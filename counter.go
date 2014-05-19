package euler_go

////////////////////////////// interface //////////////////////////////
type Counter interface {
	Next() int64
}

///////////////////////// infinite linear counter /////////////////////
type LinearCounter struct {
	Value, step int64
}

func NewLinearCounter(start, step int64) (c *LinearCounter) {
	return &LinearCounter{Value: start, step: step}
}

func (c *LinearCounter) Next() int64 {
	c.Value += c.step
	return c.Value
}

//////////////////////// infinite fibonacci series /////////////////////
type FibonacciCounter struct {
	Value, nextValue int64
}

func NewFibonacciCounter() (c *FibonacciCounter) {
	return &FibonacciCounter{1, 1}
}

func (c *FibonacciCounter) Next() int64 {
	c.Value, c.nextValue = c.nextValue, c.Value+c.nextValue
	return c.Value
}

//////////////////////////// collatz series ////////////////////////////

type CollatzSeries struct {
	Value, Length int64
}

func NewCollatzSeries(start int64) (s *CollatzSeries) {
	return &CollatzSeries{Value: start, Length: 1}
}

func (s *CollatzSeries) Next() int64 {
	if s.Value == 1 {
		return s.Value
	} else {
		if s.Value%2 == 0 {
			s.Value /= 2
			s.Length += 1
		} else {
			s.Value = s.Value*3 + 1
			s.Length += 1
		}
		return s.Value
	}
}
