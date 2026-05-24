package distributor_test

import (
	"math/rand/v2"
	"sync"
	"testing"
	"time"

	"example.com/round-robin/internal/distributor"
)

func randomDuration(maxMilli int) time.Duration {
	return time.Duration(rand.IntN(maxMilli)) * time.Millisecond
}

func TestRoundRobinDistributor(t *testing.T) {
	// Arrange
	const workers int = 5
	const iterations int = 100
	n := 10
	d := distributor.RoundRobinDistributor{N: n}

	var mu sync.Mutex
	results := make(map[int]int)

	// Act
	var wg sync.WaitGroup
	for range workers {
		wg.Go(func() {
			for range iterations {
				time.Sleep(randomDuration(10))
				v := d.Next()
				mu.Lock()
				results[v]++
				mu.Unlock()
			}
		})
	}

	wg.Wait()

	// Assert
	for i := range n {
		x, ok := results[i]
		if !ok {
			t.Errorf("Did not find value %d in results", i)
		}

		expectedFreq := (workers * iterations) / n
		if x != expectedFreq {
			t.Errorf("Expected %d occurences of %d, but found %d", expectedFreq, i, x)
		}
	}
}
