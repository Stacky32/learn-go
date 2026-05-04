package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"math/rand/v2"
	"os"
	"time"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			slog.Error("recovered from panic", "error", r)
		}
	}()

	ch := fanIn(listenOnStdIn(), listenNoise())
	timeout := time.After(20 * time.Second)

	for {
		select {
		case m := <-ch:
			slog.Info("received", "message", m)
		case <-timeout:
			slog.Info("conversation expired")
			return
		}
	}
}

func fanIn(ch1, ch2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case m1 := <-ch1:
				c <- m1
			case m2 := <-ch2:
				c <- m2
			}
		}
	}()

	return c
}

// listenOnStdIn returns a channel with strings read from standard input
func listenOnStdIn() <-chan string {
	ch := make(chan string)
	scanner := bufio.NewScanner(os.Stdin)

	go func() {
		for scanner.Scan() {
			ch <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}()

	return ch
}

// listenNoise returns a channel with strings sent at random intervals
func listenNoise() <-chan string {
	ch := make(chan string)

	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("RANDOM MESSAGE #%d", i)
			time.Sleep(time.Duration(rand.IntN(1e4)) * time.Millisecond)
		}
	}()

	return ch
}
