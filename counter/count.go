package counter

type Counter struct {
	begin, end int
}

// New create a new generator.
func New(begin, end int) *Counter {
	return &Counter{begin, end}
}

// Next returns the next value and a boolean indicating if the generator is done.
func (g *Counter) Next() (current int, done bool) {
	if done = g.begin >= g.end; done {
		return g.begin, done
	}

	current, g.begin = g.begin, g.begin+1
	return current, done
}
