package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("it can be incremented", func(t *testing.T) {
		counter := Counter{}
		expected := 3

		for i := 0; i < expected; i++ {
			counter.Inc()
		}

		assertCount(t, expected, &counter)
	})

	t.Run("it can be used concurrently", func(t *testing.T) {
		counter := Counter{}
		expected := 1000

		var wg sync.WaitGroup
		wg.Add(expected)

		for i := 0; i < expected; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCount(t, expected, &counter)
	})
}

func assertCount(t testing.TB, expected int, counter *Counter) {
	t.Helper()
	if counter.Value() != expected {
		t.Errorf("expected %d, got %d", expected, counter.Value())
	}
}
