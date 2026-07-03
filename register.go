package xk6_entry

import (
	"github.com/b0kar/xk6-counter/counter"
	"go.k6.io/k6/v2/js/modules"
)

type CounterModule struct{}

func (m *CounterModule) Up() int {
	return counter.Up()
}

func (m *CounterModule) Get() int {
	return counter.Get()
}

func (m *CounterModule) Reset() {
	counter.Reset()
}

func init() {
	modules.Register("k6/x/counter", &CounterModule{})
}
