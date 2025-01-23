package internal

import (
	"3-1/internal/contract"
)

type ChannelSubscriber struct {
	Name   string
	events map[string]func(*interface{})
}

func NewObserver(name string) contract.Observer {
	return &ChannelSubscriber{
		Name:   name,
		events: make(map[string]func(*interface{})),
	}
}

func (c *ChannelSubscriber) Update(event string, data *interface{}) {
	f, ok := c.events[event]
	if ok {
		f(data)
	}
}

func (c *ChannelSubscriber) Observe(event string, f func(data *interface{})) {
	c.events[event] = f
}
