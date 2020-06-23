package balancer

import (
	"math/rand"
	"sync"
	"time"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
)

// Intn implements rand.Intn in a concurrent-safe way.
func Intn(n int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	mu := sync.Mutex{}

	mu.Lock()
	res := r.Intn(n)
	mu.Unlock()
	return res
}

// Name is the name of our balancer.
const Name = "my_awesome_balancer"

// NewBuilder creates a new awesome balancer builder.
func NewBuilder() balancer.Builder {
	return base.NewBalancerBuilder(Name, &awPickerBuilder{}, base.Config{HealthCheck: true})
}

type awPickerBuilder struct{}

func (*awPickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	if len(info.ReadySCs) == 0 {
		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
	}
	var scs []balancer.SubConn
	for sc := range info.ReadySCs {
		scs = append(scs, sc)
	}

	return &awPicker{
		subConns: scs,
		// Start at a random index, as the same RR balancer rebuilds a new
		// picker when SubConn states change, and avoid apply excess
		// load to the first server in the list.
		next: Intn(len(scs)),
	}
}

type awPicker struct {
	subConns []balancer.SubConn
	mu       sync.Mutex
	next     int
}

func (p *awPicker) Pick(balancer.PickInfo) (balancer.PickResult, error) {
	p.mu.Lock()

	// Get conexions in order.
	sc := p.subConns[p.next]
	p.next++
	if p.next == len(p.subConns) {
		p.next = 0
	}

	p.mu.Unlock()

	return balancer.PickResult{SubConn: sc}, nil
}
