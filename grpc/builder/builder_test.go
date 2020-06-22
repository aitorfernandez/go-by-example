package builder_test

import (
	"sync"
	"testing"

	"github.com/aitorfernandez/go-by-example/grpc/builder"
	"google.golang.org/grpc/resolver"
)

type testClientConn struct {
	resolver.ClientConn
	target           string
	m1               sync.Mutex
	state            resolver.State
	updateStateCalls int
}

func (t *testClientConn) UpdateState(s resolver.State) {
	t.m1.Lock()
	defer t.m1.Unlock()
	t.state = s
	t.updateStateCalls++
}

func (t *testClientConn) getState() (resolver.State, int) {
	t.m1.Lock()
	defer t.m1.Unlock()
	return t.state, t.updateStateCalls
}

func TestBuilder(t *testing.T) {
	target := resolver.Target{
		Scheme:   "test",
		Endpoint: "grpc.io",
	}
	addrs := []string{"localhost:5001", "localhost:5002"}

	b := builder.New(target.Scheme, addrs)
	cc := &testClientConn{target: b.SchemeName}
	r, err := b.Build(target, cc, resolver.BuildOptions{})

	if err != nil {
		t.Errorf("expecting resolver %+v", r)
	}
}
