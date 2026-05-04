package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand/v2"
	"runtime"
	"sync"
	"time"
)

func main() {
	c := runtime.GOMAXPROCS(0) * 8
	xs := make([]float64, 0, c)
	for range c {
		xs = append(xs, rand.Float64()*2-1)
	}

	t := time.Now()
	if e := calculateRoots(xs); e != nil {
		log.Print(e)
	}

	log.Printf("Completed in %s", time.Since(t))
}

func calculateRoot(x float64) error {
	if x < 0 {
		return fmt.Errorf("Sqrt(%f): square root of a negative number is not real", x)
	}

	// Simulate long running calculation
	time.Sleep(time.Second * 1)

	log.Printf("Sqrt(%f)=%f", x, math.Sqrt(x))

	return nil
}

func calculateRoots(xs []float64) error {
	var mu sync.Mutex
	var errs error

	maxProcs := runtime.GOMAXPROCS(0)
	log.Printf("MAXPROCS=%d", maxProcs)
	pool := make(chan bool, maxProcs)
	var wg sync.WaitGroup
	wg.Add(len(xs))

	for _, x := range xs {
		go func(x float64) {
			defer wg.Done()

			pool <- true
			defer func() { <-pool }()

			if err := calculateRoot(x); err != nil {
				mu.Lock()
				errs = errors.Join(errs, err)
				mu.Unlock()
			}
		}(x)
	}

	wg.Wait()
	return errs
}
