package xk6_entry

import (
	"github.com/zagaris/xk6-counter/counter"
	"go.k6.io/k6/v2/js/modules"
)

type CounterModule struct{}

func (m *CounterModule) up() int {
	return counter.up()
}

func (m *CounterModule) get() int {
	return counter.get()
}

func (m *CounterModule) reset() {
	counter.reset()
}

func init() {
	modules.Register("k6/x/counter", &CounterModule{})
}
