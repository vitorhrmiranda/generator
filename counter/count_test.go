package counter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vitorhrmiranda/generator/counter"
)

func TestGenerator_Next(t *testing.T) {
	g := counter.New(0, 3)

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
}
