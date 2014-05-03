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
	return &LinearCounter{Value:start, step:step}
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
	c.Value, c.nextValue = c.nextValue, c.Value + c.nextValue
	return c.Value
}