package ranger

type Ranger struct {
	begin, end int
	c          chan int
}

// New create a new generator.
func New(begin, end int) *Ranger {
	r := &Ranger{begin, end, make(chan int, 1000)}
	go r.start()
	return r
}

// Next returns the next value and a boolean indicating if the generator is done.
func (g *Ranger) Next() (int, bool) {
	select {
	case current, ok := <-g.c:
		if !ok {
			return g.end, true
		}

		return current, false
	}
}

func (g *Ranger) start() {
	defer close(g.c)
	for i := g.begin; i < g.end; i++ {
		g.c <- i
	}
}
