package fibonacci

import "sync"

func FibRecursive(n int) int64 {
	if n == 0 || n == 1 {
		return 1
	}

	return FibRecursive(n-1) + FibRecursive(n-2)
}

type ConcurrentCache struct {
	mu sync.Mutex
	v  map[int]int64
}

func (c *ConcurrentCache) Value(k int) (val int64, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok = c.v[k]
	return val, ok
}

func (c *ConcurrentCache) Set(k int, v int64) int64 {
	c.mu.Lock()
	c.v[k] = v

	defer c.mu.Unlock()
	return v
}

func FibRecursiveCached(n int) int64 {
	cache := ConcurrentCache{v: make(map[int]int64, n+1)}
	return fibRecursiveCached(n, &cache)
}

func fibRecursiveCached(n int, c *ConcurrentCache) int64 {
	if n == 0 || n == 1 {
		return 1
	}

	if val, ok := c.Value(n); ok {
		return val
	}

	return c.Set(n-1, fibRecursiveCached(n-1, c)) + c.Set(n-2, fibRecursiveCached(n-2, c))
}
