package distributor

import "sync"

type RoundRobinDistributor struct {
	N     int
	mu    sync.Mutex
	state int
}

func (d *RoundRobinDistributor) Next() int {
	d.mu.Lock()
	defer d.mu.Unlock()

	s := (d.state + 1) % d.N
	d.state = s

	return s
}
