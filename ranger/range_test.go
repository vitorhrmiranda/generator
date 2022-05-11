package ranger_test

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vitorhrmiranda/generator/ranger"
)

func TestGenerator_Next(t *testing.T) {
	g := ranger.New(0, 3)
	time.Sleep(1 * time.Second)

	t.Run("when count until 2", func(t *testing.T) {
		var result = make([]int, 0)
		for c, done := g.Next(); !done; c, done = g.Next() {
			result = append(result, c)
		}

		assert.Equal(t, []int{0, 1, 2}, result)
	})

	t.Run("when call after done", func(t *testing.T) {
		c, done := g.Next()

		assert.Equal(t, 3, c)
		assert.Equal(t, true, done)
	})

	t.Run("when has concurrency", func(t *testing.T) {
		a1, an, n := 1, 101, 3
		g := ranger.New(a1, an)
		canal := make(chan int, (an - a1))

		var wg sync.WaitGroup

		for i := 0; i < n; i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()
				for v, done := g.Next(); !done; v, done = g.Next() {
					canal <- v
				}
			}()
		}

		wg.Wait()
		close(canal)

		var result int
		for v := range canal {
			result += v
		}

		assert.Equal(t, (an-a1)*(a1+an-1)/2, result)
	})
}
