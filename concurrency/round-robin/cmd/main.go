package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"

	"example.com/round-robin/internal/distributor"
)

func main() {
	d := distributor.RoundRobinDistributor{N: 5}
	const workers int = 5
	var wg sync.WaitGroup
	for i := range workers {
		wg.Go(func() {
			for range 10 {
				time.Sleep(randomDuration(50))
				fmt.Printf("worker %d: next = %d\n", i, d.Next())
			}
		})
	}

	wg.Wait()
}

func randomDuration(maxMilli int) time.Duration {
	return time.Duration(rand.IntN(maxMilli)) * time.Millisecond
}
